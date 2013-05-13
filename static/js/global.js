/* 导航栏高亮 */
/* 边栏菜单高亮 */
$(function() {
    var location = $('input[name=location]').val();
    $('li.' + location).addClass('active');
});

// 没有提示信息时(包括空格和空行)，不显示提示信息框
$(function() {
    if ($('div.alert span').text() == false) {
        $('div.alert').hide();
    }
});
