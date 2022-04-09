import Vue from "vue";

const colors = [
    "#F8B26A", "#F47E60", "#3498DB", "#F26D6D",
    "#67CC86", "#849B87", "#FB9B5F", "#E15B64",
    "#ABBD81"
]

function randomColor() {
    const r = Math.floor(Math.random() * colors.length);
    return colors[r]
}

Vue.prototype.$randomColor = randomColor