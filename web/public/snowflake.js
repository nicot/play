const canvas = document.getElementById('canvas');
const ctx = canvas.getContext('2d');
const image = document.getElementById('header-image')

const snowflakes = [];

function drawSnowflake(e) {
    ctx.fillStyle = 'white';
    ctx.fillRect(e.offsetX, e.offsetY, 5, 5);
    console.log(e.offsetX, e.offsetX);
}

function resize() {
    canvas.width = image.clientWidth;
    canvas.height = image.clientHeight;
}
export default function () {
    canvas.addEventListener('mousedown', e => drawSnowflake(e))
    window.addEventListener('resize', resize);
    resize();
}