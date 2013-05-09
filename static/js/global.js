/* 边栏菜单高亮 */
$(function() {
    var location = $('input[name=location]').val();
    $('li.' + location).addClass('active');
});
