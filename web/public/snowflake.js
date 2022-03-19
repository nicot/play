const canvas = document.getElementById('canvas');
const ctx = canvas.getContext('2d');
const image = document.getElementById('header-image')

const snowflakes = [];
let mouseUpSnowedAt = new Date();
let mousePos = null;
let mouseDown = false;
let randomSnowedAt = new Date();

function addSnowflakes() {
    const now = Date.now();
    if (mouseDown) {
        snowflakes.push({ x: mousePos.x, y: mousePos.y })
        mouseUpSnowedAt = now;
    } else if (mousePos && now - mouseUpSnowedAt > 200) {
        snowflakes.push({ x: mousePos.x, y: mousePos.y })
        mouseUpSnowedAt = now;
    }
    if (now - randomSnowedAt > 1000) {
        snowflakes.push({
            x: Math.random() * canvas.width,
            y: Math.random() * canvas.height,
        });
        randomSnowedAt = now;
    }
}

function resize() {
    canvas.width = image.clientWidth;
    canvas.height = image.clientHeight;
}

function draw() {
    addSnowflakes()

    ctx.clearRect(0, 0, canvas.width, canvas.height);

    snowflakes.forEach(snowflake => {
        snowflake.y += .6;
    })

    snowflakes.forEach(snowflake => {
        ctx.beginPath();
        ctx.strokeStyle = 'white';
        ctx.arc(snowflake.x, snowflake.y, 3, 0, Math.PI * 2);
        ctx.closePath();
        ctx.stroke();
    });

    window.requestAnimationFrame(draw)
}

export default function () {
    canvas.addEventListener('mousemove', e =>
        mousePos = { x: e.offsetX, y: e.offsetY })
    canvas.addEventListener('mousedown', () => mouseDown = true)
    canvas.addEventListener('mouseup', () => mouseDown = false)

    window.addEventListener('resize', resize);
    window.requestAnimationFrame(draw)
    resize();
}