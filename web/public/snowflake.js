const canvas = document.getElementById('canvas');
const ctx = canvas.getContext('2d');
const image = document.getElementById('header-image')

const snowflakes = [];
let mouseUpSnowedAt = new Date();
let mousePos = null;
let mouseDown = false;
let randomSnowedAt = new Date();

function newSnowflake(x, y) {
    const direction = 1 - Math.round(Math.random()) * 2;
    snowflakes.push({ x, y, direction, showPercent: 100 })
}

function addSnowflakes() {
    const now = Date.now();
    if (mouseDown) {
        newSnowflake(mousePos.x, mousePos.y);
        mouseUpSnowedAt = now;
    } else if (mousePos && now - mouseUpSnowedAt > 200) {
        newSnowflake(mousePos.x, mousePos.y);
        mouseUpSnowedAt = now;
    }
    if (now - randomSnowedAt > 1000) {
        const x = (Math.random() + .1) / 1.2 * canvas.width;
        const y = Math.random() * image.clientHeight * .8;
        newSnowflake(x, y);
        randomSnowedAt = now;
    }
}

function resize() {
    canvas.width = document.body.clientWidth;
    canvas.height = document.body.clientHeight;
}

function drawSnowflake({ x, y, showPercent }) {
    ctx.save();

    ctx.translate(x, y);
    ctx.beginPath();
    const show = showPercent / 100 * .6;
    ctx.strokeStyle = `rgba(255, 255, 255, ${show})`;
    ctx.moveTo(0, 0);

    for (let i = 0; i < 6; i++) {
        ctx.lineWidth = 2
        ctx.moveTo(0, 0);
        ctx.lineTo(3, 3);
        ctx.rotate(Math.PI / 3);
    }
    ctx.stroke();

    ctx.restore();
}

function draw() {
    addSnowflakes()

    ctx.clearRect(0, 0, canvas.width, canvas.height);

    snowflakes.forEach(snowflake => {
        snowflake.y += .6;
        snowflake.x += .3 * snowflake.direction * Math.random();
        if (Math.random() < 1 / 500) {
            snowflake.direction *= -1;
        }
    })

    for (let i = 0; i < snowflakes.length; i++) {
        const past = snowflakes[i].y - image.clientHeight;
        const fadeOut = 100;
        if (past > 0) {
            snowflakes[i].showPercent = (fadeOut - past) / fadeOut * 100;
        }
        if (past > fadeOut) {
            snowflakes.splice(i, 1);
        }
    }

    snowflakes.forEach(drawSnowflake);

    window.requestAnimationFrame(draw);
}

export default function () {
    canvas.addEventListener('mousemove', e =>
        mousePos = { x: e.offsetX, y: e.offsetY })
    canvas.addEventListener('mousedown', e => {
        mousePos = { x: e.offsetX, y: e.offsetY };
        mouseDown = true;
    })
    canvas.addEventListener('mouseup', () => mouseDown = false)

    window.addEventListener('resize', resize);
    window.addEventListener('load', resize);
    window.requestAnimationFrame(draw)
    resize();
}
