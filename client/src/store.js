export default {
    state: {
        screen: '',
        playing: false,
        categoriesPlaying: [],
    },
    mutations: {
        startGame: (state, cb = function() {}) => {
            state.playing = true
            cb(state)
        },
        stopGame: (state, cb = function() {}) => {
            state.playing = false
            state.categoriesPlaying = []
            cb(state)
        },
        setCategories: (state, cats, cb = function() {}) => {
            state.categoriesPlaying = cats
            cb(state)
        },
        setScreen: (state, screen, cb = function() {}) => {
            state.screen = screen
            cb(state)
        },
    },
}
