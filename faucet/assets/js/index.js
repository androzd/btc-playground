$(document).on('click', '[data-ajax-href]', (event) => {
    event.preventDefault();
    let $target = $(event.currentTarget);
    let href = $target.data('ajax-href');
    let method = $target.data('ajax-method') || 'get';
    let data = $target.data()
    $.ajax({
        url: href,
        method: method,
        success: (result) => {
            console.log(result.status)
        }
    });
    console.log(href, method);
})

let setBadge = (status) => {
    let $badge = $('#status-badge');
    $badge.text(status);
    if (status == 'working') {
        $badge.removeClass('badge-danger').addClass('badge-success');
    } else {
        $badge.removeClass('badge-success').addClass('badge-danger');
    }
}
let updateBadge = () => {
    $.ajax({
        url: '/api/mining/status',
        method: 'GET',
        success: (result) => {
            setBadge(result.status)
        }
    })
}
setInterval(() => {
    updateBadge()
}, 1000);
updateBadge()