<template>
  <div style="display: flex" >
      <div style="display: flex">
        <div style="flex-direction: column">
          <div id="main" style="width: 600px; height: 400px;"></div>
          <div id="main2" style="width: 600px; height: 400px;"></div>
        </div>
      </div>
      <div style="display: flex">
        <div style="flex-direction: column">
          <div id="main3" style="width: 600px; height: 400px;"></div>
          <div id="main4" style="width: 600px; height: 400px;"></div>
        </div>
    </div>
  </div>
</template>

<script>
import * as echarts from 'echarts';
import axios from "axios";
var myChart;
var myChart2;
var myChart3;
var myChart4;
export default {
  data(){
    return{
      university:[],
      grateCounts:[]
    }
  },
  unmounted(){
    this.university = []
    this.grateCounts = []

  },
  async created() {
    await axios.get('http://localhost:8080/colleges')
        .then(response => {
          this.university = response.data;
        })
        .catch(error => {
    });
    const gradeCounts = {
      "A+":  this.university.filter(item => item.grade === "A+").length,
      "A":  this.university.filter(item => item.grade === "A").length,
      "A-":  this.university.filter(item => item.grade === "A-").length,
      "B+":  this.university.filter(item => item.grade === "B+").length,
      "B":  this.university.filter(item => item.grade === "B").length,
      "B-":  this.university.filter(item => item.grade === "B-").length,
    };
    // 基于准备好的 dom，初始化 echarts 实例
    if (myChart != null && myChart != "" && myChart != undefined) {
      myChart.dispose(); //销毁
    }
    if (myChart2 != null && myChart2 != "" && myChart2 != undefined) {
      myChart2.dispose(); //销毁
    }
    if (myChart3 != null && myChart3 != "" && myChart3 != undefined) {
      myChart3.dispose(); //销毁
    }
    if (myChart4 != null && myChart4 != "" && myChart4 != undefined) {
      myChart4.dispose(); //销毁
    }
    myChart = echarts.init(document.getElementById('main'));
    myChart2 = echarts.init(document.getElementById('main2'));
    myChart3 = echarts.init(document.getElementById('main3'));
    myChart4 = echarts.init(document.getElementById('main4'));
    // 指定图表的配置项和数据
    const option = {
      title: {
        text: '计算机科学与技术专业大学排行榜'
      },
      tooltip: {},
      legend: {
        data: ['学校数']
      },
      xAxis: {
        data: ['A+','A','A-','B+','B','B-']
      },
      yAxis: {},
      series: [
        {
          name: '等级',
          type: 'bar',
          data: [gradeCounts["A+"].toString(), gradeCounts["A"].toString(), gradeCounts["A-"].toString(),
            gradeCounts["B+"].toString(), gradeCounts["B"].toString(), gradeCounts["B-"].toString()]
        }
      ]
    };

    // 使用刚指定的配置项和数据显示图表。
    myChart.setOption(option);

    const option2 = {
      tooltip: {
        trigger: 'item'
      },
      legend: {
        top: '5%',
        left: 'center'
      },
      series: [
        {
          name: 'Access From',
          type: 'pie',
          radius: ['40%', '70%'],
          avoidLabelOverlap: false,
          itemStyle: {
            borderRadius: 10,
            borderColor: '#fff',
            borderWidth: 2
          },
          label: {
            show: false,
            position: 'center'
          },
          emphasis: {
            label: {
              show: true,
              fontSize: 40,
              fontWeight: 'bold'
            }
          },
          labelLine: {
            show: false
          },
          data: [
            { value: gradeCounts["A+"], name:  "A+"},
            { value: gradeCounts.A, name: "A" },
            { value: gradeCounts["A-"], name: "A-" },
            { value: gradeCounts["B+"], name: "B+" },
            { value: gradeCounts.B, name: "B" },
            { value: gradeCounts["B-"], name: "B-"}
          ]
        }
      ]
    };

    myChart2.setOption(option2)

    const option3 = {
      xAxis: {
        type: 'category',
        data: ['A+', 'A', 'A-', 'B+', 'B', 'B-']
      },
      yAxis: {
        type: 'value'
      },
      series: [
        {
          data: [gradeCounts["A+"],gradeCounts.A,gradeCounts["A-"],gradeCounts["B+"],gradeCounts.B,gradeCounts["B-"]],
          type: 'line'
        }
      ]
    };
    myChart3.setOption(option3);

    const option4 = {
      xAxis: {
        type: 'category',
        data:  ['A+', 'A', 'A-', 'B+', 'B', 'B-']
      },
      yAxis: {
        type: 'value'
      },
      series: [
        {
          data: [gradeCounts["A+"],gradeCounts.A,gradeCounts["A-"],gradeCounts["B+"],gradeCounts.B,gradeCounts["B-"]],
          type: 'line',
          smooth: true
        }
      ]
    };
    myChart4.setOption(option4);
  }
};
</script>
