$(function() {
    var row_width = $('body').width();
    var sum_width = 0;
    $('nav.site-nav > ul > li').each(function(){
        sum_width += $(this).width();
    });

    //TODO: this is just a hack now.
    var ratio = row_width / sum_width;
    if(ratio < 1) {
        var original_font = $('nav.site-nav').css('font-size').replace('px', '');
        var ratio_font = parseFloat(original_font) * ratio;
        $('nav.site-nav').css('font-size', Math.floor(ratio_font));
    }
});
