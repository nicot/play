const { FuseBox } = require("fuse-box");
const fuse = FuseBox.init({
    homeDir: "./src",
    output: "./public/out/$name.js",
})
fuse.bundle("app")
    .instructions(`>index.ts`)
    .hmr()
    .watch()
fuse.dev()
fuse.run()