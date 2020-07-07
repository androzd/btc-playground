(function ($) {
    $.fn.serializeFormJSON = function () {
        var o = {};
        var a = this.serializeArray();
        var form = this;
        $.each(a, function () {
            if ('number' == $(form).find('[name='+this.name+']').attr('type')) {
                this.value = parseFloat(this.value);
            }
            if (o[this.name]) {
                if (!o[this.name].push) {
                    o[this.name] = [o[this.name]];
                }
                o[this.name].push(this.value || '');
            } else {
                o[this.name] = this.value || '';
            }
        });
        return o;
    };
})(jQuery);

$(document).on('click', '[data-ajax-href]', (event) => {
    event.preventDefault()
    event.stopPropagation()
    let $target = $(event.currentTarget);
    let href = $target.data('ajax-href');
    let method = $target.data('ajax-method') || 'get';
    let data = $target.data()
    $.ajax({
        url: href,
        method: method,
        success: (result) => {
            setBadge(result.status)
        }
    });
})

$(document).on('submit', '[data-ajax-form]', (event) => {
    event.preventDefault()
    let $form = $(event.currentTarget);
    let href = $form.attr('action');
    let method = $form.attr('method') || 'get';
    let data = $form.serializeFormJSON();
    console.log({url: href,
        method: method,
        contentType: "application/json",
        dataType: "json",
        data: data});
    $.ajax({
        url: href,
        method: method,
        contentType: "application/json",
        data: JSON.stringify(data),
        success: (result) => {
            console.log(result)
        }
    });
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