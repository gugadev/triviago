<style scoped>
    .loading {
        color: rgba(255,255,255,.6);
        font-family: 'Pacifico', sans-serif;
        font-size: 30px;
        font-weight: 400;
        margin-bottom: 40px;
        text-align: center;
    }
    .container {
        align-items: center;
        display: flex;
        flex-direction: column;
    }
    .screen {
        margin: 50px 0 0 0;
        position: relative;
    }
    .question {
        padding: 25px 40px;
        background-color: #f4f4f4;
        background-color: rgba(0,0,0,.4);
        margin: 0 auto;
        position: relative;
        max-width: 90%;
    }
    .question .cat-pic {
        background-repeat: no-repeat;
        background-size: cover;
        height: 50px;
        left: -15px;
        position: absolute;
        top: -15px;
        transform: rotate(-45deg);
        width: 50px;
    }
    .question p {
        font-family: 'Luckiest Guy', sans-serif;
        color: #555;
        color: rgba(255,255,255,.8);
        font-weight: 600;
        text-align: center;
    }
    .options {
        padding: 25px;
    }
    .options ul {
        display: flex;
        flex-wrap: wrap;
        justify-content: space-between;
        list-style: none;
    }
    .option {
        border-radius: 25px;
        cursor: pointer;
        font-family: 'Luckiest Guy', sans-serif;
        font-size: 14px;
        flex: 0 0 45%;
        margin: 6px 0;
        padding: 12px 15px;
        text-align: center;
        background-color: rgba(0,0,0,.25);
        border: 2px solid rgba(255,255,255,.6);
        color: rgba(255,255,255,.8);
        font-weight: 600;
        transition: all .25s ease;
    }
    .option:hover {
        border-color: #fff;
        color: #fff;
    }
    .option.selected {
        background-color: rgba(20, 100, 180, .4);
        border-color: rgba(255,255,255,.9);
        color: #fff;
    }
    .option.hit {
        background-color: rgba(20, 180, 40, .4);
    }
    .option.wrong {
        background-color: rgba(180, 40, 20, .4);
    }
    .counter {
        color: rgba(255,255,255,.8);
        font-family: 'Luckiest guy', sans-serif;
        font-size: 38px;
        position:absolute;
        right: 20px;
        top: -50px;
    }
</style>

<template>
    <section class="container">
        <transition name="fade">
            <!-- Loading indicator -->
            <h2 class="loading" v-show="loading">Cargando...</h2>
        </transition>
        <transition name="fade">
            <!-- Game self -->
            <section class="screen" v-if="ready && playing">
                <countdown
                        :time="15000"
                        :auto-start="false"
                        :emit-events="true"
                        @countdownend="onTimeout"
                        ref="counter"
                        v-if="!answer.id"
                >
                    <template slot-scope="props">
                        <span
                                class="counter"
                                :style="{ 'color': props.seconds < 10 ? '#d53400' : 'rgba(255,255,255,.8)'}"
                        >
                            {{props.seconds}}
                        </span>
                    </template>
                </countdown>
                <section class="question">
                    <img :src="icons[question.category_id]" alt="cat" class="cat-pic">
                    <p>{{question.text}}</p>
                </section>
                <section class="options">
                    <ul>
                        <li
                                class="option"
                                :class="{'selected': answer.id === op.id, 'hit': (answer.id === op.id) && success, 'wrong': (answer.id === op.id) && !success}"
                                v-for="op in options"
                                @click="choose(op)"
                        >
                            {{op.text}}
                        </li>
                    </ul>
                </section>
            </section>
        </transition>

    </section>
</template>

<script>
    import Axios from 'axios'
    import Categories from '../components/categories.vue'
    import icons from '../icons.json'

    export default {
        components: {
            Categories,
        },
        computed: {
            playing() { return this.$store.state.playing },
            categories() { return this.$store.state.categoriesPlaying },
        },
        data: () => ({
            question: {
                category_id: null,
                id: null,
                text: null,
            },
            options: [],
            answer: {},
            ready: false,
            loading: true,
            success: false,
            icons,
        }),
        methods: {
            choose(option) {
                this.answer = option
                this.check()
            },
            start() {
                this.ask().then((data) => {
                    this.loading = true
                    this.question = data.question
                    this.options = data.answers
                    setTimeout(() => {
                        this.loading = false
                        setTimeout(() => {
                            this.ready = true
                            this.$store.commit('startGame')
                            this.$nextTick(() => {
                                this.$refs.counter.start()
                            })
                        }, 800)
                    }, 2500)
                })
            },
            ask() {
                let query = 'cats='
                for(let cat of this.categories) {
                    query += `${cat.id},`
                }
                query = query.substr(0, query.length - 1)
                return (
                    Axios
                        .get(`/api/questions?${query}`)
                        .then(({ data }) => data)
                )
            },
            next() {
                this.ask().then((data) => {
                    this.answer = {}
                    this.$nextTick(() => {
                        this.question = data.question
                        this.options = data.answers
                        this.success = false
                        this.$refs.counter.start()
                    })
                })
            },
            onTimeout() {
                this.answer = this.options.find(o => o.correct)
                this.success = true
                this.next()
            },
            check() {
                Axios
                    .post(`/api/answers/check?question=${this.question.id}&answer=${this.answer.id}`)
                    .then(({ data }) => {
                        this.success = data.success
                        setTimeout(() => {
                            this.next()
                        }, 1000)
                    })
            },
        },
        mounted() {
            this.$store.commit('setScreen', 'game')
            this.start()
        },
    }
</script>