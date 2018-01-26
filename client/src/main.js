import Vue from 'vue/dist/vue'
import Vuex from 'vuex'
import VueRouter from 'vue-router'
import VueChart from 'vue-chartjs'
import CountDown from '@xkeshi/vue-countdown'
import Container from './components/container.vue'
import Home from './screens/home.vue'
import Game from './screens/game.vue'
import Score from './screens/score.vue'
import Create from './screens/create.vue'
import NotFound from './screens/404.vue'
import storeConfig from './store'


Vue.use(VueRouter)
Vue.use(Vuex)
Vue.use(VueChart)
Vue.component('countdown', CountDown)

const store = new Vuex.Store(storeConfig)

const router = new VueRouter({
    mode: 'history',
    routes: [
        { path: '/', name: 'home', component: Home },
        { path: '/game', name: 'game', component: Game },
        { path: '/create', name: 'create', component: Create },
        { path: '/score', name: 'score', component: Score },
        { path: '*', component: NotFound },
    ],
})

router.beforeEach((from, to, next) => {
    if (from.path === '/game' && !store.state.playing) {
        next('/')
    } else {
        next()
    }
})

const app = new Vue({
    components: { Container },
    template: '<Container />',
    router,
    store,
})

app.$mount('#root')
