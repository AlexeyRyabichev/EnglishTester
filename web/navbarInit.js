document.addEventListener('DOMContentLoaded', function() {
    var elems = document.querySelectorAll('.sidenav');
    var instances = M.Sidenav.init(elems, {});

    var elemsDropdown = document.querySelectorAll('.dropdown-trigger');
    var instancesDropdown = M.Dropdown.init(elemsDropdown, {
        coverTrigger: false,
        hover: true,
        constrainWidth: false
    });

    var elemsCollapsible = document.querySelectorAll('.collapsible');
    var instancesCollapsible = M.Collapsible.init(elemsCollapsible, {});
});