<style scoped>
h1 {
    color: rgba(255,255,255,.8);
    font-family: 'Luckiest Guy';
    font-size: 28px;
    margin-top: 10px;
    padding: 10px;
    text-align: center;
}
input, select, textarea {
    border: none;
    border-radius: 8px;
    color: #444;
    font-family: 'Luckiest Guy';
    height: 35px;
    padding: 0 15px;
}
textarea {
    height: auto;
    padding: 12px 15px;
}
.question, .answers {
    align-items: flex-start;
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    padding: 10px;
}
.question textarea,
.question select,
.answers input {
    flex: 0 0 47.5%;
    font-size: 14px;
    width: 0;
}
.answers input,
.asnwers textarea {
    margin: 5px 0;
}
input.correct {
    color: #20bf6b;
}
input.correct::placeholder {
    color: rgba(32, 191, 107, .7);
}
.create-btn {
    border: none;
    bottom: 30px;
    left: 50%;
    transition: all .1s ease;
    margin-top: 15px;
    padding: 10px 40px;
    border-radius: 10px;
    cursor: pointer;
    font-family: 'Pacifico', cursive;
    font-size: 18px;
    color: #FFF;
    text-decoration: none;
    transform: translateX(-50%);
    background-color: #1abc9c;
    border-bottom: 5px solid #16a085;
    outline: none;
    text-shadow: 0px -2px #16a085;
    position: absolute;
    white-space: nowrap;
}
.create-btn:disabled {
    background-color: #f9f9f9;
    color: #d6d6d6;
    text-shadow: 0 -2px #d6d6d6;
    border-bottom-color: #d6d6d6;
    cursor: not-allowed;
    pointer-events: none;
}
.create-btn:not(:disabled):active {
    transform: translate(-50%, 5px);
    border-bottom: 1px solid #2980B9;
}
</style>

<template>
    <form @submit.prevent="create">
        <header>
            <h1>¡Crea tu pregunta!</h1>
        </header>
        <div class="question">
            <textarea
                rows="5"
                placeholder="Escribe tu pregunta"
                v-model="question.text"
                @input="validate"
            />
            <select
                v-model="question.category_id"
                @change="validate"
                required>
                <option :value="undefined">Categoría..</option>
                <option
                    v-for="cat in categories"
                    v-bind:key="cat.id"
                    :value="cat.id"
                >
                    {{cat.name}}
                </option>
            </select>
        </div>
        <div class="answers">
            <input
                type="text"
                class="correct"
                placeholder="Respuesta correcta"
                v-model="answers[0].text"
                @input="validate"
                required
            >
            <input
                type="text"
                placeholder="Respuesta.."
                v-model="answers[1].text"
                @input="validate"
                required
            >
            <input
                type="text"
                placeholder="Respuesta.."
                v-model="answers[2].text"
                @input="validate"
                required
            >
            <input
                type="text"
                placeholder="Respuesta.."
                v-model="answers[3].text"
                @input="validate"
                required
            >
        </div>
        <button
            :disabled="!valid"
            class="create-btn"
            v-if="screen === 'create'"
        >
            <span>Crear Pregunta</span>
        </button>
  </form>
</template>

<script>
import Axios from 'axios'

export default {
    computed: {
        screen() {
            return this.$store.state.screen
        },
    },
    data: () => ({
        categories: [],
        question: { text: '', category_id: undefined },
        answers: [
            { text: '', correct: true },
            { text: '', correct: false },
            { text: '', correct: false },
            { text: '', correct: false },
        ],
        valid: false,
    }),
    methods: {
        create() {
            this.shuffleAnswers()
            Axios
            .post('/api/questions', {
                question: this.question,
                answers: this.answers,
            })
            .then(({ data }) => {
                if (data === 'success') {
                    alert("Pregunta creada")
                    this.reset()
                }
            })
        },
        reset() {
            this.question = { text: '', category_id: undefined }
            this.answers = [
                { text: '', correct: true },
                { text: '', correct: false },
                { text: '', correct: false },
                { text: '', correct: false },
            ]
        },
        validate() {
            const questionValid = this.question.text !== '' && this.question.category_id !== undefined
            const answersValid = !this.answers.find(a => a.text === '')
            this.valid = questionValid && answersValid
        },
        shuffleAnswers() {
            const answers = [...this.answers]
            let counter = answers.length;

            while (counter > 0) {
                // Pick a random index
                const index = Math.floor(Math.random() * counter);
                // Decrease counter by 1
                counter--;
                // And swap the last element with it
                const temp = answers[counter];
                answers[counter] = answers[index];
                answers[index] = temp;
            }
            this.answers = answers
        },
    },
    mounted() {
        this.$store.commit('setScreen', 'create')
        Axios.get("/api/categories")
                .then(({ data }) => {
                    this.categories = data
                })
    },
}
</script>

