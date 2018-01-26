<style scoped>
    .loading {
        color: rgba(255,255,255,.6);
        font-family: 'Pacifico', sans-serif;
        font-size: 30px;
        font-weight: 400;
        margin-bottom: 40px;
        text-align: center;
    }
    .score-container {
        padding: 20px;
    }
    .global-score {
        display: flex;
        justify-content: space-between;
    }
    .score-data {
        display: flex;
        flex-direction: column;
        justify-content: center;
        padding: 15px 10px;
    }
    .score-data span {
        font-family: 'Luckiest Guy', sans-serif;
        font-size: 72px;
    }
    .score-data label {
        color: rgba(255,255,255,.85);
        font-family: 'Luckiest Guy', sans-serif;
        margin-top: 12px;
        font-size: 18px;
    }
    .score-data.total span {
        color: #f39c12;
    }
    .score-data.hits span {
        color: #27ae60;
    }
</style>

<template>
    <section class="score-container">
        <transition name="fade">
            <h1 class="loading" v-if="loading">Cargando...</h1>
        </transition>
        <transition name="fade">
            <section class="global-score" v-if="ready">
                <div class="score-data total">
                    <span>{{total}}</span>
                    <label>Respondidas</label>
                </div>
                <div class="score-data hits">
                    <span>{{hits}}</span>
                    <label>Correctas</label>
                </div>
            </section>
        </transition>
        <transition name="fade">
            <Chart :labels="labels" :data="data" v-if="ready" :height="300" />
        </transition>
    </section>
</template>

<script>
    import Axios from 'axios'
    import Chart from '../components/scorechart.vue'

    export default {
        components: {
            Chart,
        },
        data: () => ({
            labels: [],
            data: [],
            total: 0,
            hits: 0,
            loading: true,
            ready: false,
        }),
        methods: {
            getStats() {
                Axios.get('/api/score').then(({ data }) => {
                    this.labels = data.map(s => s.category)
                    this.data = data.map(s => s.hits)
                    this.total = data.reduce((p, c) => p + c.total, 0)
                    this.hits = data.reduce((p, c) => p + c.hits, 0)
                    this.loading = false
                    setTimeout(() => {
                        this.ready = true
                    }, 700)
                })
            },
        },
        mounted() {
            this.$store.commit('setScreen', 'score')
            this.getStats()
        }
    }
</script>