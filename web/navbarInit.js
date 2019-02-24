const M = require("./js/bin/materialize");
document.addEventListener('DOMContentLoaded', function () {
    let elements = document.querySelectorAll('.sidenav');
    // noinspection JSUnusedLocalSymbols
    let instances = M.Sidenav.init(elements, {});
});