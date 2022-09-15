
function trim (str) { //去掉头尾空格
  return str.replace(/(^\s*)|(\s*$)/g, "");
}

function cx (token, num) {
  $('#submit').hide();


  $('#load').html('正在通关,请稍等...');
  var i = 0;
  var win = 0;
  var count = 0;
  for (i; i < num; i++) {
    $.ajax({
      type: "POST",
      url: "http://gg.liujiaweixiaoman.cn/chabai/v1/",
      contentType: "application/json",
      data: JSON.stringify({ "token": token }),
      dataType: "json",
      success: function (data) {
        if (data.code === 200) {
          $('#log').html($('#log').html() + "第" + (count+1) + "次成功" + '\n');
          win++;
        } else {
          $('#log').html($('#log').html() + "第" + (count+1) + "次失败" + '\n' + "已结束!");
          return;
        }
        count++;
        if (count == num) {
          $('#submit').show();
          $('#load').html('已通关' + win + '次');

        }

      },

    });



    //$('#submit').show();
  }
}
$(document).ready(function () {

  $('#submit').click(function () {
    var self = $(this);
    var token = trim($('#uin').val());
    var num = trim($('#num').val());
    if (token == '') {
      alert("请输入Token！");
      return false;
    }

    $('#load').show();
    $("#log").html("");
    cx(token, num);
  });


});