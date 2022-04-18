<template>
    <div class="charts">
        <div class="pie-chart" id="pie-chart"></div>
    </div>
</template>

<script>
export default {
    name: "BlogTypePie",
    data() {
      return {
      }
    },
    methods: {
        createPie(val) {
            const chartDom = document.getElementById('pie-chart');
            const myChart = this.$echarts.init(chartDom);
            myChart.setOption(val)
        },
        async getData() {
            var options = {
                title: {
                    text: '博客分类统计',
                        left: 'center'
                },
                tooltip: {
                    trigger: 'item'
                },
                series: [
                    {
                        name: '分类名称',
                        type: 'pie',
                        radius: '50%',
                        data: [
                            { value: 1048, name: '博客系统' },
                            { value: 735, name: 'Golang' },
                            { value: 580, name: 'Mysql' },
                            { value: 484, name: 'Redis' },
                            { value: 300, name: 'Nginx' },
                        ],
                        emphasis: {
                            itemStyle: {
                                shadowBlur: 10,
                                shadowOffsetX: 0,
                                shadowColor: 'rgba(0, 0, 0, 0.5)'
                            }
                        }
                    }
                ]
            }
            const {data:res} = await this.$axios.get("/myblog/staticsBlog")
            if (res.status !== 1) {
                this.$message.warning("获取博客信息失败!")
                return
            }

            if (res.data.length <= 0) {
                return
            }

            const arr = res.data[0].map(({count, name}) => {
                return {value: count, name: name}
            })
            options.series[0].data = [...arr]
            this.createPie(options)
        }
    },
    mounted() {
        this.getData()
    },
}
</script>

<style scoped>
.charts {
    position: relative;
    height: 320px;
    width: 300px;
    background-color: #fff;
    border-radius: 12px;
    padding: 18px 0;
}

.pie-chart {
    position: absolute;
    width: 300px;
    height: 300px;
    left: 50%;
    transform: translateX(-50%);
}

</style>
