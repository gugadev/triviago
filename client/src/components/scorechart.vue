<script>
    import { Doughnut } from 'vue-chartjs'

    export default {
        extends: Doughnut,
        props: ['labels', 'data'],
        methods: {
            buildChart() {
                Chart.defaults.global.elements.arc.borderWidth = 0
                const data = {
                    labels: this.labels,
                    datasets: [
                        {
                            data: this.data,
                            backgroundColor: [
                                "#FF6384",
                                "#36A2EB",
                                "#FFCE56",
                                "#1abc9c",
                                "#8e44ad",
                                "#7f8c8d",
                                "#d35400",
                                "#e74c3c",
                            ],
                        },
                    ],
                }
                this.renderChart(data, {
                    responsive: true,
                    legend: {
                        labels: {
                            fontColor: 'rgba(255,255,255,.85)',
                            fontFamily: 'Luckiest Guy',
                            fontSize: 13,
                        },
                    },
                    elements: {
                        arc: {
                            borderWidth: 0,
                        },
                    },
                })
            },
        },
        mounted() {
            this.buildChart()
        },
    }
</script>
