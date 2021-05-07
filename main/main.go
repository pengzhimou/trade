package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
	"trade/tradecenter/sdktrade"

	"github.com/stretchr/testify/assert"
	chart "github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

const (
	lineChartXAxisName  = "Date"
	lineChartYAxisName  = "Count"
	lineChartHeight     = 700
	lineChartWidth      = 1280
	colorMultiplier     = 256
	imgStrPrefix        = "data:image/png;base64,"
	pieLabelFormat      = "%v %v"
	barChartTryAgainErr = "invalid data range; cannot be zero"
)

var (
	lineChartStyle = chart.Style{
		Padding: chart.Box{
			Top:  30,
			Left: 150,
		},
	}

	defaultChartStyle = chart.Style{
		Padding: chart.Box{
			Top: 30,
		},
	}

	timeFormat = chart.TimeDateValueFormatter
)

type LineYValue struct {
	Name   string
	Values []float64
}

type ChartValue struct {
	Name  string
	Value float64
}

// createLineChart 创建线性图
func createLineChart(title string, endTime time.Time, values []LineYValue) (img string, err error) {
	if len(values) == 0 {
		return
	}
	// 1、计算X轴
	lenX := len(values[0].Values)
	// X轴内容xValues 及 X轴坐标ticks
	var xValues []time.Time
	var ticks []chart.Tick
	for i := lenX - 1; i >= 0; i-- {
		curTime := endTime.AddDate(0, 0, -i)
		xValues = append(xValues, curTime)
		ticks = append(ticks, chart.Tick{Value: getNsec(curTime), Label: timeFormat(curTime)})
	}

	// 2、生成Series
	var series []chart.Series
	for _, yValue := range values {
		series = append(series, chart.TimeSeries{
			Name: yValue.Name,
			Style: chart.Style{
				// 随机渲染线条颜色
				StrokeColor: drawing.Color{
					R: uint8(rand.Intn(colorMultiplier)),
					G: uint8(rand.Intn(colorMultiplier)),
					B: uint8(rand.Intn(colorMultiplier)),
					A: uint8(colorMultiplier - 1), // 透明度
				},
			},
			XValues: xValues,
			YValues: yValue.Values,
		})
	}

	// 3、新建图形
	graph := chart.Chart{
		Title:      title,
		Background: lineChartStyle,
		Width:      lineChartWidth,
		Height:     lineChartHeight,
		XAxis: chart.XAxis{
			Name:           lineChartXAxisName,
			ValueFormatter: timeFormat,
			Ticks:          ticks,
		},
		YAxis: chart.YAxis{
			Name: lineChartYAxisName,
		},
		Series: series,
	}
	graph.Elements = []chart.Renderable{
		chart.LegendLeft(&graph),
	}

	// 4、输出目标
	img, err = writeLineChart(&graph)

	return
}

// getNsec 获取纳秒数
func getNsec(cur time.Time) float64 {
	return float64(cur.Unix() * int64(time.Second))
}

func writeLineChartToPng(c *chart.Chart) (img string, err error) {
	f, _ := os.Create("graph.png")
	err = c.Render(chart.PNG, f)
	return
}

func writeLineChart(c *chart.Chart) (img string, err error) {
	var imgContent bytes.Buffer
	err = c.Render(chart.PNG, &imgContent)
	if err != nil {
		return
	}

	img = imgToStr(imgContent)
	return
}

func imgToStr(imgContent bytes.Buffer) string {
	return imgStrPrefix + base64.StdEncoding.EncodeToString(imgContent.Bytes())
}

// createPieChart 创建饼图
func createPieChart(title string, pieValues []ChartValue) (img string, err error) {
	if len(pieValues) == 0 {
		return
	}
	// 1、构建value
	var values []chart.Value
	for _, v := range pieValues {

		values = append(values, chart.Value{
			Value: v.Value,
			// Label: fmt.Sprintf(pieLabelFormat, getSimpleSensType(v.Name), formatValue(v.Value)),
			Label: fmt.Sprintf(pieLabelFormat, v.Name, formatValue(v.Value)),
		})
	}

	// 2、新建饼图
	pie := chart.PieChart{
		Title:      title,
		Background: defaultChartStyle,
		Values:     values,
	}

	// 4、输出目标
	img, err = writePieChart(&pie)

	return
}

func formatValue(f float64) string {
	return fmt.Sprintf("%.2fW", f/10000)
}

func writePieChartToPng(c *chart.PieChart) (img string, err error) {
	f, _ := os.Create("pie.png")
	err = c.Render(chart.PNG, f)
	return
}

func writePieChart(c *chart.PieChart) (img string, err error) {
	var imgContent bytes.Buffer
	err = c.Render(chart.PNG, &imgContent)
	if err != nil {
		return
	}

	img = imgToStr(imgContent)
	return
}

// createBarChart 创建柱状图
func createBarChart(title string, barValues []ChartValue) (img string, err error) {
	if len(barValues) == 0 {
		return
	}
	// 1、构建value
	var values []chart.Value
	for _, v := range barValues {
		values = append(values, chart.Value{
			Value: v.Value,
			Label: v.Name,
		})
	}

	// 2、新建饼图
	bar := chart.BarChart{
		XAxis: chart.Style{
			TextWrap: 0, // default 1为可以溢出规定的范围
		},
		Width:      2560,
		BarWidth:   50,
		BarSpacing: 300,
		Title:      title,
		Background: defaultChartStyle,
		Bars:       values,
	}

	// 4、输出目标
	img, err = writeBarChart(&bar)
	if err != nil && err.Error() == barChartTryAgainErr {
		// 添加一个隐藏条目，设置透明度A为0, 设置任意属性如R不为0即可
		values = append(values, chart.Value{
			Style: chart.Style{
				StrokeColor: drawing.Color{R: 1},
			},
			Value: 0,
			Label: "",
		})
		bar.Bars = values
		img, err = writeBarChart(&bar)
	}

	return
}

func writeBarChartToPng(c *chart.BarChart) (img string, err error) {
	f, _ := os.Create("bar.png")
	err = c.Render(chart.PNG, f)
	return
}

func writeBarChart(c *chart.BarChart) (img string, err error) {
	var imgContent bytes.Buffer
	err = c.Render(chart.PNG, &imgContent)
	if err != nil {
		return
	}

	img = imgToStr(imgContent)
	return
}

func TestCreateLineChart(t *testing.T) {
	testAssert := assert.New(t)

	tests := []struct {
		title     string
		endTime   time.Time
		barValues []LineYValue
	}{
		{"line chart", time.Now(), []LineYValue{
			{"asd", []float64{1, 2, 300, 100, 200, 6, 700}},
			{"hgj", []float64{400, 500000, 200, 50, 5, 800, 7}},
			{"dfg45r", []float64{1, 2, 700, 100, 200, 6, 700}},
			{"2342sr", []float64{400, 500000, 200, 50, 5, 800, 7}},
			{"das21-asd", []float64{300000, 200000, 400000, 100000, 400000, 450000, 400000}},
			{"csc", []float64{400, 500000, 200, 50, 5, 800, 7}},
			{"mhj", []float64{1, 2, 300, 100, 200, 6, 700}},
			{"876ijgh", []float64{400, 500000, 200, 50, 5, 800, 7}},
			{"fbfdv", []float64{1, 2, 300, 100, 200, 6, 700}},
			{"67ds", []float64{400, 10000, 200, 50, 5, 800, 7}},
			{"67bdfv", []float64{1, 2, 300, 100, 200, 6, 700}},
			{"sdf324", []float64{400, 500000, 200, 50, 5, 800, 7}},
			{"vdf67", []float64{1, 2, 300, 100, 200, 6, 700}},
			{"vdfs234", []float64{400, 500000, 200, 50, 5, 800, 7}},
			{"123sdf", []float64{1, 2, 700, 100, 200, 6, 700}},
			{"aasdasd", []float64{400, 500000, 200, 50, 5, 800, 7}},
			{"aasd", []float64{1, 2, 300, 100, 200, 6, 700}},
			{"basd", []float64{400, 500000, 200, 50, 5, 800, 7}},
			{"cczx", []float64{1, 2, 300, 100, 200, 6, 700}},
			{"qweqw", []float64{400, 500000, 200, 50, 5, 800, 7}},
			{"asdadf", []float64{1, 2, 300, 100, 200, 6, 700}},
			{"fghfh", []float64{400, 500000, 200, 50, 5, 800, 7}},
			{"erttyrt", []float64{1, 2, 300, 100, 200, 6, 700}}}},
	}

	for _, test := range tests {
		img, err := createLineChart(test.title, test.endTime, test.barValues)
		testAssert.Equal(img, "")
		testAssert.Equal(err, nil)
	}
}

func TestCreatePieChart(t *testing.T) {
	testAssert := assert.New(t)

	tests := []struct {
		title     string
		pieValues []ChartValue
	}{
		{"pie chart", []ChartValue{{"asdas", 20000}, {"q12asd", 300000}, {"ascasd", 3000}}},
	}

	for _, test := range tests {
		img, err := createPieChart(test.title, test.pieValues)
		testAssert.Equal(img, "")
		testAssert.Equal(err, nil)
	}
}

func TestCreateBarChart(t *testing.T) {
	testAssert := assert.New(t)

	tests := []struct {
		title     string
		pieValues []ChartValue
	}{
		{"bar chart", []ChartValue{{"asdascasd\nasd-asd", 20}, {"asdascascasdasdasd.go\nasdasd-asdasd", 30}, {"asasdasd.asdasd]\nasdasd-asda", 100},
			{"asdasdasda.go\nasdasd-asdasd", 20}, {"asdasd.asdasd\ngeass", 30}, {"asdasdasd\nasdasd-asdasd", 100},
			{"asdasd_adsdasd_dasd\asd-asd", 20}, {"asdascasdcad\nasdasdasda", 30}, {"asdasdasdasd", 100},
			{"asdasclkhy9p867p9", 20}}},
	}

	for _, test := range tests {
		img, err := createBarChart(test.title, test.pieValues)
		testAssert.Equal(img, "")
		testAssert.Equal(err, nil)
	}
}

func main() {

	// sdktrade.TradeTest3()
	// sdktrade.TradeTest2()

	sdktrade.RsiTest()

}
