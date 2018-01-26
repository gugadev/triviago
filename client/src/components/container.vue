<style scoped>
    .container {
        align-items: center;
        display: flex;
        flex-direction: column;
        height: 100vh;
        padding: 50px 0;
        margin: 0 auto;
        max-width: 100%;
        width: 500px;
    }
    .game-board {
        max-width: 100%;
        width: 500px;
    }
    .logo {
        margin-top: 10px;
        max-width: 90%;
    }
    .start-btn, .stop-btn {
        border: none;
        bottom: 30px;
        transition: all .1s ease;
        margin-top: 15px;
        padding: 10px 40px;
        border-radius: 10px;
        cursor: pointer;
        font-family: 'Pacifico', cursive;
        font-size: 18px;
        color: #FFF;
        text-decoration: none;
        background-color: #1abc9c;
        border-bottom: 5px solid #16a085;
        outline: none;
        text-shadow: 0px -2px #16a085;
        position: absolute;
    }
    .stop-btn {
        background-color: #f39c12;
        border-bottom-color: #DB8D0F;
        text-shadow: 0px -2px #DB8D0F;
    }
    .start-btn:not(:disabled):active {
        transform: translate(0px,5px);
        border-bottom: 1px solid #2980B9;
    }
    .stop-btn:not(:disabled):active {
        border-bottom: 1px solid #DB8D0F;
    }
    .start-btn:disabled, .stop-btn:disabled {
        background-color: #f9f9f9;
        color: #d6d6d6;
        text-shadow: 0 -2px #d6d6d6;
        border-bottom-color: #d6d6d6;
        cursor: not-allowed;
    }
    .score-btn, .create-btn, .home-btn {
        background-color: transparent;
        border: none;
        cursor: pointer;
        outline: none;
        position: fixed;
        right: 10px;
        top: 10px;
        width: 40px;
    }
    .home-btn {
        left: 10px;
        right: auto;
    }
    .create-btn {
        left: 50%;
        transform: translateX(-50%);
    }
</style>

<template>
    <main class="container">
        <img src="/static/img/logo.png" alt="logo" class="logo">
        <button class="home-btn" @click="goHome">
            <img src="/static/img/start.svg" alt="score">
        </button>
        <button class="create-btn" @click="goCreate">
            <img src="/static/img/create.svg" alt="create">
        </button>
        <button class="score-btn" @click="goScore">
            <img src="/static/img/score.svg" alt="score">
        </button>
        <article class="game-board">
            <transition name="screenfade">
                <router-view></router-view>
            </transition>
        </article>
        <button
            :disabled="!categories.length"
            @click="startGame"
            class="start-btn"
            v-if="screen === 'home'"
        >
            <span>Iniciar juego</span>
        </button>
        <button
            :disabled="!categories.length"
            @click="stopGame"
            class="stop-btn"
            v-if="screen === 'game'"
        >
            <span>Parar juego</span>
        </button>
    </main>
</template>

<script>
    export default {
        computed: {
            screen() {
                return this.$store.state.screen
            },
            playing() {
                return  this.$store.state.playing
            },
            categories() {
                return this.$store.state.categoriesPlaying
            },
        },
        methods: {
            startGame() {
                this.$store.commit('startGame', () => {
                    console.log('Iniciando juego...')
                    this.$router.push({ name: 'game' })
                })
            },
            stopGame() {
                this.$store.commit('stopGame', () => {
                    console.log('Parando juego...')
                    this.$router.push({ name: 'home' })
                })
            },
            goHome() {
                if (this.screen === 'game') {
                    return
                }
                this.$router.push({ name: 'home' })
            },
            goCreate() {
                this.$router.push({ name: 'create' })
            },
            goScore() {
                this.$router.push({ name: 'score' })
            },
        },
    }
</script>