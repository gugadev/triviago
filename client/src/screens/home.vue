<template>
    <Categories />
</template>
<script>
    import Categories from '../components/categories.vue'

    export default {
        components: {
            Categories,
        },
        mounted() {
            this.$store.commit('setScreen', 'home')
        },
    }
</script>
