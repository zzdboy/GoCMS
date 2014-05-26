var ColorHex=new Array('00','33','66','99','CC','FF')
var SpColorHex=new Array('FF0000','00FF00','0000FF','FFFF00','00FFFF','FF00FF')
var current=null
var colorTable=''

function colorpicker(showid,fun) {
	for (i=0;i<2;i++) {
	  for (j=0;j<6;j++) {
		colorTable=colorTable+'<tr height=12>'
		colorTable=colorTable+'<td width=11 onmouseover="onmouseover_color(\'000\')" onclick="select_color(\''+showid+'\',\'000\','+fun+')" style="background-color:#000">'
		if (i==0){
		colorTable=colorTable+'<td width=11 onmouseover="onmouseover_color(\''+ColorHex[j]+ColorHex[j]+ColorHex[j]+'\')" onclick="select_color(\''+showid+'\',\''+ColorHex[j]+ColorHex[j]+ColorHex[j]+'\','+fun+')" style="background-color:#'+ColorHex[j]+ColorHex[j]+ColorHex[j]+'">'
		}  else {
		colorTable=colorTable+'<td width=11 onmouseover="onmouseover_color(\''+SpColorHex[j]+'\')" onclick="select_color(\''+showid+'\',\''+SpColorHex[j]+'\','+fun+')" style="background-color:#'+SpColorHex[j]+'">'} 
	
		colorTable=colorTable+'<td width=11 onmouseover="onmouseover_color(\'000\')" onclick="select_color(\''+showid+'\',\'000\','+fun+')" style="background-color:#000">'
		for (k=0;k<3;k++) {
		   for (l=0;l<6;l++) {
			colorTable=colorTable+'<td width=11 onmouseover="onmouseover_color(\''+ColorHex[k+i*3]+ColorHex[l]+ColorHex[j]+'\')" onclick="select_color(\''+showid+'\',\''+ColorHex[k+i*3]+ColorHex[l]+ColorHex[j]+'\','+fun+')"  style="background-color:#'+ColorHex[k+i*3]+ColorHex[l]+ColorHex[j]+'">'
		   }
		 }
	  }
	}
	colorTable='<div style="position:relative;width:253px; height:176px"><a href="javascript:;" onclick="closeBox();" class="close-own">X</a><table width=253 border="0" cellspacing="0" cellpadding="0" style="border:1px #000 solid;border-bottom:none;border-collapse: collapse" bordercolor="000000">'
			   +'<tr height=30><td colspan=21 bgcolor=#eeeeee>'
			   +'<table cellpadding="0" cellspacing="1" border="0" style="border-collapse: collapse">'
			   +'<tr><td width="3"><td><input type="text" name="DisColor" size="6" id="background_colorId" disabled style="border:solid 1px #000000;background-color:#ffff00"></td>'
			   +'<td width="3"><td><input type="text" name="HexColor" size="7" id="input_colorId" style="border:inset 1px;font-family:Arial;" value="#000000"></td><td><a href="javascript:;" onclick="clear_title();"> clear</a></td></tr></table></td></table>'
			   +'<table border="1" cellspacing="0" cellpadding="0" style="border-collapse: collapse" bordercolor="000000" style="cursor:hand;">'
			   +colorTable+'</table></div>';
	$('#'+showid).html(colorTable);
	colorTable = '';
}
function onmouseover_color(color) {
	var color = '#'+color;
	$('#background_colorId').css('background-color',color);
	$('#input_colorId').val(color);
  
}
function select_color(showid,color,fun) {
	var color = '#'+color;
	//$('#title').css('color',color);
	if(fun) {
		fun.apply(this,[color]);
	}
	$('#'+showid).html(' ');
}
function closeBox(){
	$(".colorpanel").html(' ');
}
function clear_title() {
	$('#title').css('color','');
	$('#title_colorpanel').html(' ');
}