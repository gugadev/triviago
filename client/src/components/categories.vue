<style scoped>
    .game-categories {
        width: 100%;
    }
    .loading {
        color: rgba(255,255,255,.8);
        font-family: 'Pacifico', sans-serif;
        font-size: 30px;
        font-weight: 400;
        margin-bottom: 40px;
        text-align: center;
    }
    .title {
        color: rgba(255,255,255,.8);
        font-family: 'Pacifico', sans-serif;
        font-size: 32px;
        font-weight: 600;
        padding: 0 15px 30px 15px;
        text-align: center;
    }
    .categories {
        display: flex;
        flex-wrap: wrap;
        justify-content: space-between;
        list-style: none;
        padding: 0 20px;
        width: 100%;
    }
    .cat {
        background-color: rgba(0,0,0,.25);
        border: 2px solid rgba(255,255,255,.7);
        border-radius: 25px;
        cursor: pointer;
        margin: 10px 0;
        padding: 12px;
        position: relative;
        flex: 0 0 45%;
        text-align: center;
        transition: all .25s ease;
    }
    .cat:hover {
        border-color: #fff;
    }
    .cat:hover span {
        color: #fff;
    }
    .cat.selected {
        background-color: rgba(20, 100, 180, .4);
        border-color: rgba(255,255,255,1);
    }
    .cat.selected span {
        color: rgba(255,255,255,1);
    }
    .cat span {
        color: rgba(255,255,255,.85);
        font-family: 'Luckiest Guy', sans-serif;
        font-size: 14px;
        font-weight: 600;
        transition: color .25s ease;
    }
    .cat img {
        height: 30px;
        left: 10px;
        position: absolute;
        top: 7px;
        width: 30px;
    }
</style>

<template>
    <div class="game-categories">
        <transition name="fade">
            <h1 class="loading" v-if="loading">Cargando...</h1>
            <h1 class="title" v-if="!loading">¿Cuáles quieres jugar?</h1>
        </transition>
        <transition name="fade">
            <ul class="categories" v-show="categories.length">
                <li
                        class="cat"
                        :class="{'selected': isSelected(null, c)}"
                        v-for="c in categories"
                        @click="select(c)"
                >
                    <img :src="icons[c.id]" />
                    <span>{{ c.name }}</span>
                </li>
            </ul>
        </transition>
    </div>
</template>

<script>
    import Axios from 'axios'
    import store from '../store'
    import icons from '../icons.json'

    export default {
        computed: {
          selected() {
              return this.$store.state.categoriesPlaying
          }
        },
        data: () => ({
            loading: true,
            categories: [],
            icons,
        }),
        methods: {
            select(c) {
                let selected = [...this.selected]
                if (!this.isSelected(selected, c)) {
                    selected.push(c)
                } else {
                    selected = selected.filter(cat => cat.id !== c.id)
                }
                this.$store.commit('setCategories', selected)
            },
            isSelected(cats, cat) {
                return (cats || this.selected).some(c => c.id === cat.id)
            },
        },
        mounted() {
            Axios.get("/api/categories").then(({ data }) => {
                setTimeout(() => {
                    this.loading = false
                }, 1250)
                setTimeout(() => {
                    this.categories = data
                }, 1500)
            })
        },
    }
</script>