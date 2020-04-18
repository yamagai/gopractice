$(function () {
  var text = $('#times').text();
  console.log(text);
  var delyear = text.replace("2020-", "");
  console.log(delyear);
  var delT = delyear.replace("T", " ");
   console.log(delT);
 var delyear2 = delT.replace("2020-", "");
 console.log(delyear2);
 var delT2 = delyear2.replace("T", " ");
  console.log(delT2);

});
