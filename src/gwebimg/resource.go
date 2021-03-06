package main

var Index_menu_html=string(`
<!doctype html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">

	<title>{{.Title}}</title>

	<link rel="stylesheet" type="text/css" href="/static/include/pure-min.css">
	<link rel="stylesheet" type="text/css" href="/static/include/side-menu.css">
	<!-- include scooch.css -->
	<link rel="stylesheet" type="text/css" href="/static/include/scooch.min.css">
	<link rel="stylesheet" type="text/css" href="/static/include/scooch-style.min.css">
	<style>
		.menu-item-selected {
			background: #333;
		}
		.menu-item-deselected {
			background: #777;
		}
	</style>
</head>
<body>
<div id="layout">
	<!-- Menu toggle -->
	<a href="#menu" id="menuLink" class="menu-link">
		<!-- Hamburger icon -->
		<span></span>
	</a>

	<div id="menu">
		<div class="pure-menu pure-menu-scrollable">
			<a class="pure-menu-heading" href="#">Command</a>
			<ul class="pure-menu-list">
				<li class="pure-menu-item"><a href="#" class="pure-menu-link" id="cmdPrev">Prev</a></li>
				<li class="pure-menu-item"><a href="#" class="pure-menu-link" id="cmdNext">Next</a></li>
				<li class="pure-menu-item"><a href="#" class="pure-menu-link" id="cmdLast">Last view</a></li>
				<li class="pure-menu-heading">Chapter list</li>
				{{range .ChapterList}}
				<li class="pure-menu-item chapter-name"><a href="#" class="pure-menu-link">{{.}}</a></li>
				{{end}}
			</ul>
		</div>
	</div>

	<div id="main">
		<div class="content">
			<div class="pure-g">
				<div class="pure-u-1">
					<!-- the viewport -->
					<div class="m-scooch m-fluid m-scooch-photos m-center">
						<!-- the slider -->
						<div class="m-scooch-inner">
							<!-- the items -->
						</div>
						<!-- the controls -->
						<div class="m-scooch-controls m-scooch-hud">
							<a href="#" data-m-slide="prev">Previous</a>
							<a href="#" data-m-slide="next">Next</a>
					</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
<!-- include zepto.js or jquery.js -->
<script src="/static/include/zepto.min.js" type="application/javascript"></script>
<!-- include scooch.js -->
<script src="/static/include/scooch.min.js" type="application/javascript"></script>
<script src="/static/include/ui.js"></script>
<!-- construct the carousel -->
<script type="application/javascript">
var page_title="{{.Title}}";
var force_load_flag = true;
var last_chapter = getCookie("chapter");
var last_page = Number(getCookie("page"));
var flag_restore_last_page = false;

console.log("last chapter:", last_chapter);
console.log("last page:", last_page);

function jump_last_view() {
	
	if (last_page == NaN) 
		last_page = 1;
	flag_restore_last_page = true;
	loadchapter(last_chapter);
	$('.pure-menu-item').each(function(i,e) { 
		if (last_chapter == $(e).text() ) {
			select_menuitem(e);
		}
	});
}

function setCookie(key, value) {
	var expires = new Date();
	expires.setTime(expires.getTime() + (1 * 24 * 60 * 60 * 1000));
	document.cookie = key + '=' + value + ';expires=' + expires.toUTCString();
}

function getCookie(key) {
	var keyValue = document.cookie.match('(^|;) ?' + key + '=([^;]*)(;|$)');
	return keyValue ? keyValue[2] : null;
}

function load_img_to_scooch(img_array) {
	if ((img_array == null) || (img_array.length == 0))
		return;
	force_load_flag = true;
	$(".m-scooch-inner").empty();
	
	for (index=0; index<img_array.length; index++) {
		var is_active="";
		if (index==0) is_active="m-active" ;
		$(".m-scooch-inner").append("<div class=\"m-item "+is_active+"\"><img src=\""+ img_array[index] +"\"/></div>");
	}
	
	$(".m-scooch").scooch("refresh");
	if (flag_restore_last_page) {
		$(".m-scooch").scooch("move",last_page);
	} else {
		$(".m-scooch").scooch("move",1);
	}
	force_load_flag = false;
	flag_restore_last_page = false;
	console.log("finish load img to scooch");
}

function load_chapterlist () {
	//$('a.chapter-name').click(function(e) { alert($(this).text()); });
	var first_chapter = $('.chapter-name').first().text();
	loadchapter(first_chapter);
	select_menuitem($('.chapter-name').first());
}

function loadchapter(chapter_name) {
	var retval=[];
	$.ajax({url:"/page-list/"+chapter_name, success: function(result) {
		console.log("finish getting pages, load to scooch");
		load_img_to_scooch(result);
		setCookie("chapter", chapter_name);
	}});
}

function page_init() {
	load_chapterlist();
	console.log("finish page init");
}

function select_menuitem(i) {
	$('.chapter-name').removeClass('pure-menu-selected'); 
	$(i).addClass('pure-menu-selected');
}

function prev_chapter() {
	var prev_chapter_name = $(".pure-menu-selected").prev().text();
	console.log("prev chapter name:", prev_chapter_name);
	if (prev_chapter_name == "Chapter List") {
		return;
	} else {
		var prev_item = $(".pure-menu-selected").prev();
		loadchapter(prev_chapter_name);
		select_menuitem(prev_item);
	}
}

function next_chapter() {
	var next_chapter_name = $(".pure-menu-selected").next().text();
	console.log("next chapter name:", next_chapter_name);
	if (next_chapter_name == null) {
		return;
	} else {
		var next_item = $(".pure-menu-selected").next();
		loadchapter(next_chapter_name);
		select_menuitem(next_item);
	}
}

page_init();   
$('.m-scooch').scooch();

$("#cmdPrev").click(function() {
	prev_chapter();
});

$("#cmdNext").click(function() {
	next_chapter();
});

$('.m-scooch').on('afterSlide', function (e, index, newIndex) {
	// find current chapter pages
	console.log(index,"-",newIndex);
	setCookie("page",newIndex);
	var img_array = $(".m-item>img")
	if (index==newIndex) {
		if (index==1) {
			if (!force_load_flag) {
				console.log("move to prev chapter");
				prev_chapter();
			}
		} else {
			console.log("move to next chapter");
			next_chapter();
		}
	} else {
		$(window).scrollTop(0);
	}
	var new_title="" + page_title + " - " + $(".pure-menu-selected").text() + " - " + newIndex + "/" + $('.m-item').length ; 
	$("title").text(new_title);
});

$("#cmdLast").click(function(){
	jump_last_view();
});

$(".chapter-name").click(function(e){
	var clicked_chapter_name = $(this).text();
	loadchapter(clicked_chapter_name);
	select_menuitem(this);
});

$(document).keydown(function(e) {
	switch(e.which) {
		case 33: //pgdown
			prev_chapter();
		break;
		case 34: //pgdown
			next_chapter();
		break;
		case 37: // left
			$('.m-scooch').scooch('prev');
		break;

		case 39: // right
			$('.m-scooch').scooch('next');
		break;

		default: return; // exit this handler for other keys
	}
	e.preventDefault(); // prevent the default action (scroll / move caret)
});

</script>

</body>
</html>
`)

var Index_html=string(`
<!doctype html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    <title>{{.Title}}</title>

    <link rel="stylesheet" type="text/css" href="/static/include/pure-min.css">
<!-- include scooch.css -->
    <link rel="stylesheet" type="text/css" href="/static/include/scooch.min.css">
    <link rel="stylesheet" type="text/css" href="/static/include/scooch-style.min.css">
    <style>
    .small-font {
    	font-size: 80%;
    }
    .button-success,
    .button-secondary {
            color: white;
            border-radius: 4px;
            text-shadow: 0 1px 1px rgba(0, 0, 0, 0.2);
    }
    .button-success {
            background: rgb(28, 184, 65); /* this is a green */
    }
    .button-secondary {
            background: rgb(66, 184, 221); /* this is a light blue */
    }
    </style>
</head>

<body>
    <!--
    Your HTML goes here. Visit purecss.io/layouts/ for some sample HTML code.
    -->
    <form class="pure-form small-font">
    <div class="pure-g">
      <div class="pure-u-1" align="center">
		  <span class="btn-prev"><a class="lnk-prev pure-button button-success small-font" href="#">Prev</a></span>
          <select id="top-chapter-select" class="chapter-select small-font">
          {{range .ChapterList}}
          <option value="{{.}}">{{.}}</option>
          {{end}}
          </select>
		  <span class="btn-prev"><a class="lnk-next pure-button button-success small-font" href="#">Next</a></span>          
		  <span class="btn-restore"><a id="last-view" class="pure-button button-secondary small-font" href="#">Last view</a></span>          
      </div>
      <div class="pure-u-1">
		<!-- the viewport -->
		<div class="m-scooch m-fluid m-scooch-photos m-center">
		  <!-- the slider -->
		  <div class="m-scooch-inner">
		    <!-- the items -->
		  </div>
		  <!-- the controls -->
		  <div class="m-scooch-controls m-scooch-hud">
		    <a href="#" data-m-slide="prev">Previous</a>
		    <a href="#" data-m-slide="next">Next</a>
		  </div>
		</div>
      </div>
      <div class="pure-u-1" align="center">
		  <span class="btn-prev"><a class="lnk-prev pure-button button-success small-font" href="#">Prev</a></span>
          <select id="bottom-chapter-select" class="chapter-select small-font">
          {{range .ChapterList}}
          <option value="{{.}}">{{.}}</option>
          {{end}}
          </select>
		  <span class="btn-prev"><a class="lnk-next pure-button button-success small-font" href="#">Next</a></span>          
      </div>
    </div>
	</form>
<!-- include zepto.js or jquery.js -->
<script src="/static/include/zepto.min.js" type="application/javascript"></script>
<!-- include scooch.js -->
<script src="/static/include/scooch.min.js" type="application/javascript"></script>
<!-- construct the carousel -->
<script type="application/javascript">
var page_title="{{.Title}}";
var force_load_flag = true;
var last_chapter = getCookie("chapter");
var last_page = Number(getCookie("page"));
var flag_restore_last_page = false;

console.log("last chapter:", last_chapter);
console.log("last page:", last_page);

function jump_last_view() {
	
	if (last_page == NaN) 
		last_page = 1;
	flag_restore_last_page = true;
	$(".chapter-select").val(last_chapter);
	$("#top-chapter-select").change();
}

function setCookie(key, value) {
	var expires = new Date();
	expires.setTime(expires.getTime() + (1 * 24 * 60 * 60 * 1000));
	document.cookie = key + '=' + value + ';expires=' + expires.toUTCString();
}

function getCookie(key) {
	var keyValue = document.cookie.match('(^|;) ?' + key + '=([^;]*)(;|$)');
	return keyValue ? keyValue[2] : null;
}

function load_img_to_scooch(img_array) {
	if ((img_array == null) || (img_array.length == 0))
		return;
	force_load_flag = true;
	$(".m-scooch-inner").empty();
	
	for (index=0; index<img_array.length; index++) {
		var is_active="";
		if (index==0) is_active="m-active" ;
		$(".m-scooch-inner").append("<div class=\"m-item "+is_active+"\"><img src=\""+ img_array[index] +"\"/></div>");
	}
	
	$(".m-scooch").scooch("refresh");
	if (flag_restore_last_page) {
		$(".m-scooch").scooch("move",last_page);
	} else {
		$(".m-scooch").scooch("move",1);
	}
	force_load_flag = false;
	flag_restore_last_page = false;
	console.log("finish load img to scooch");
}

function load_chapterlist () {
	var first_chapter = $('#top-chapter-select>option').first().val();
	loadchapter(first_chapter);
	
}

function loadchapter(chapter_name) {
	var retval=[];
	$.ajax({url:"/page-list/"+chapter_name, success: function(result) {
		console.log("finish getting pages, load to scooch");
		load_img_to_scooch(result);
		setCookie("chapter", chapter_name);
	}});
}

function page_init() {
	load_chapterlist();
	console.log("finish page init");
}

function prev_chapter() {
	var current_selected_chapter = $("#top-chapter-select").val();
	var next_chapter = $("#top-chapter-select>option").filter(function( index ){ return $(this).val()== current_selected_chapter }).prev() ;
	if (next_chapter.length == 1) {
		$("#top-chapter-select").val(next_chapter.val());
		$("#bottom-chapter-select").val(next_chapter.val());
		$("#top-chapter-select").change();
	}
}

function next_chapter() {
	var current_selected_chapter = $("#top-chapter-select").val();
	var next_chapter = $("#top-chapter-select>option").filter(function( index ){ return $(this).val()== current_selected_chapter }).next() ;
	if (next_chapter.length == 1) {
		$("#top-chapter-select").val(next_chapter.val());
		$("#bottom-chapter-select").val(next_chapter.val());
		$("#top-chapter-select").change();
	}
}

page_init();   
$('.m-scooch').scooch();

$(".chapter-select").change( function (e){
	//var selected_chapter = $(this).val();
	var selected_chapter = this.value;
	console.log("chapter change");
	loadchapter(selected_chapter);
	$(".chapter-select").not(this).val( $(this).val() );
	
});

$(".lnk-prev").click(function() {
	prev_chapter();
});

$(".lnk-next").click(function() {
	next_chapter();
});

$('.m-scooch').on('afterSlide', function (e, index, newIndex) {
	// find current chapter pages
	console.log(index,"-",newIndex);
	setCookie("page",newIndex);
	var img_array = $(".m-item>img")
	if (index==newIndex) {
		if (index==1) {
			if (!force_load_flag) {
				console.log("move to prev chapter");
				prev_chapter();
			}
		} else {
			console.log("move to next chapter");
			next_chapter();
		}
	} else {
		$(window).scrollTop(0);
	}
	var new_title="" + page_title + " - " + $("#top-chapter-select").val() + " - " + newIndex + "/" + $('.m-item').length ; 
	$("title").text(new_title);
});

$("#last-view").click(function(){
	jump_last_view();
});
</script>
</body>
</html>
`)

var Pure_min_css = []byte(`
/*!
Pure v0.6.0
Copyright 2014 Yahoo! Inc. All rights reserved.
Licensed under the BSD License.
https://github.com/yahoo/pure/blob/master/LICENSE.md
*/
/*!
normalize.css v^3.0 | MIT License | git.io/normalize
Copyright (c) Nicolas Gallagher and Jonathan Neal
*/
/*! normalize.css v3.0.2 | MIT License | git.io/normalize */html{font-family:sans-serif;-ms-text-size-adjust:100%;-webkit-text-size-adjust:100%}body{margin:0}article,aside,details,figcaption,figure,footer,header,hgroup,main,menu,nav,section,summary{display:block}audio,canvas,progress,video{display:inline-block;vertical-align:baseline}audio:not([controls]){display:none;height:0}[hidden],template{display:none}a{background-color:transparent}a:active,a:hover{outline:0}abbr[title]{border-bottom:1px dotted}b,strong{font-weight:700}dfn{font-style:italic}h1{font-size:2em;margin:.67em 0}mark{background:#ff0;color:#000}small{font-size:80%}sub,sup{font-size:75%;line-height:0;position:relative;vertical-align:baseline}sup{top:-.5em}sub{bottom:-.25em}img{border:0}svg:not(:root){overflow:hidden}figure{margin:1em 40px}hr{-moz-box-sizing:content-box;box-sizing:content-box;height:0}pre{overflow:auto}code,kbd,pre,samp{font-family:monospace,monospace;font-size:1em}button,input,optgroup,select,textarea{color:inherit;font:inherit;margin:0}button{overflow:visible}button,select{text-transform:none}button,html input[type=button],input[type=reset],input[type=submit]{-webkit-appearance:button;cursor:pointer}button[disabled],html input[disabled]{cursor:default}button::-moz-focus-inner,input::-moz-focus-inner{border:0;padding:0}input{line-height:normal}input[type=checkbox],input[type=radio]{box-sizing:border-box;padding:0}input[type=number]::-webkit-inner-spin-button,input[type=number]::-webkit-outer-spin-button{height:auto}input[type=search]{-webkit-appearance:textfield;-moz-box-sizing:content-box;-webkit-box-sizing:content-box;box-sizing:content-box}input[type=search]::-webkit-search-cancel-button,input[type=search]::-webkit-search-decoration{-webkit-appearance:none}fieldset{border:1px solid silver;margin:0 2px;padding:.35em .625em .75em}legend{border:0;padding:0}textarea{overflow:auto}optgroup{font-weight:700}table{border-collapse:collapse;border-spacing:0}td,th{padding:0}.hidden,[hidden]{display:none!important}.pure-img{max-width:100%;height:auto;display:block}.pure-g{letter-spacing:-.31em;*letter-spacing:normal;*word-spacing:-.43em;text-rendering:optimizespeed;font-family:FreeSans,Arimo,"Droid Sans",Helvetica,Arial,sans-serif;display:-webkit-flex;-webkit-flex-flow:row wrap;display:-ms-flexbox;-ms-flex-flow:row wrap;-ms-align-content:flex-start;-webkit-align-content:flex-start;align-content:flex-start}.opera-only :-o-prefocus,.pure-g{word-spacing:-.43em}.pure-u{display:inline-block;*display:inline;zoom:1;letter-spacing:normal;word-spacing:normal;vertical-align:top;text-rendering:auto}.pure-g [class *="pure-u"]{font-family:sans-serif}.pure-u-1,.pure-u-1-1,.pure-u-1-2,.pure-u-1-3,.pure-u-2-3,.pure-u-1-4,.pure-u-3-4,.pure-u-1-5,.pure-u-2-5,.pure-u-3-5,.pure-u-4-5,.pure-u-5-5,.pure-u-1-6,.pure-u-5-6,.pure-u-1-8,.pure-u-3-8,.pure-u-5-8,.pure-u-7-8,.pure-u-1-12,.pure-u-5-12,.pure-u-7-12,.pure-u-11-12,.pure-u-1-24,.pure-u-2-24,.pure-u-3-24,.pure-u-4-24,.pure-u-5-24,.pure-u-6-24,.pure-u-7-24,.pure-u-8-24,.pure-u-9-24,.pure-u-10-24,.pure-u-11-24,.pure-u-12-24,.pure-u-13-24,.pure-u-14-24,.pure-u-15-24,.pure-u-16-24,.pure-u-17-24,.pure-u-18-24,.pure-u-19-24,.pure-u-20-24,.pure-u-21-24,.pure-u-22-24,.pure-u-23-24,.pure-u-24-24{display:inline-block;*display:inline;zoom:1;letter-spacing:normal;word-spacing:normal;vertical-align:top;text-rendering:auto}.pure-u-1-24{width:4.1667%;*width:4.1357%}.pure-u-1-12,.pure-u-2-24{width:8.3333%;*width:8.3023%}.pure-u-1-8,.pure-u-3-24{width:12.5%;*width:12.469%}.pure-u-1-6,.pure-u-4-24{width:16.6667%;*width:16.6357%}.pure-u-1-5{width:20%;*width:19.969%}.pure-u-5-24{width:20.8333%;*width:20.8023%}.pure-u-1-4,.pure-u-6-24{width:25%;*width:24.969%}.pure-u-7-24{width:29.1667%;*width:29.1357%}.pure-u-1-3,.pure-u-8-24{width:33.3333%;*width:33.3023%}.pure-u-3-8,.pure-u-9-24{width:37.5%;*width:37.469%}.pure-u-2-5{width:40%;*width:39.969%}.pure-u-5-12,.pure-u-10-24{width:41.6667%;*width:41.6357%}.pure-u-11-24{width:45.8333%;*width:45.8023%}.pure-u-1-2,.pure-u-12-24{width:50%;*width:49.969%}.pure-u-13-24{width:54.1667%;*width:54.1357%}.pure-u-7-12,.pure-u-14-24{width:58.3333%;*width:58.3023%}.pure-u-3-5{width:60%;*width:59.969%}.pure-u-5-8,.pure-u-15-24{width:62.5%;*width:62.469%}.pure-u-2-3,.pure-u-16-24{width:66.6667%;*width:66.6357%}.pure-u-17-24{width:70.8333%;*width:70.8023%}.pure-u-3-4,.pure-u-18-24{width:75%;*width:74.969%}.pure-u-19-24{width:79.1667%;*width:79.1357%}.pure-u-4-5{width:80%;*width:79.969%}.pure-u-5-6,.pure-u-20-24{width:83.3333%;*width:83.3023%}.pure-u-7-8,.pure-u-21-24{width:87.5%;*width:87.469%}.pure-u-11-12,.pure-u-22-24{width:91.6667%;*width:91.6357%}.pure-u-23-24{width:95.8333%;*width:95.8023%}.pure-u-1,.pure-u-1-1,.pure-u-5-5,.pure-u-24-24{width:100%}.pure-button{display:inline-block;zoom:1;line-height:normal;white-space:nowrap;vertical-align:middle;text-align:center;cursor:pointer;-webkit-user-drag:none;-webkit-user-select:none;-moz-user-select:none;-ms-user-select:none;user-select:none;-webkit-box-sizing:border-box;-moz-box-sizing:border-box;box-sizing:border-box}.pure-button::-moz-focus-inner{padding:0;border:0}.pure-button{font-family:inherit;font-size:100%;padding:.5em 1em;color:#444;color:rgba(0,0,0,.8);border:1px solid #999;border:0 rgba(0,0,0,0);background-color:#E6E6E6;text-decoration:none;border-radius:2px}.pure-button-hover,.pure-button:hover,.pure-button:focus{filter:progid:DXImageTransform.Microsoft.gradient(startColorstr='#00000000', endColorstr='#1a000000', GradientType=0);background-image:-webkit-gradient(linear,0 0,0 100%,from(transparent),color-stop(40%,rgba(0,0,0,.05)),to(rgba(0,0,0,.1)));background-image:-webkit-linear-gradient(transparent,rgba(0,0,0,.05) 40%,rgba(0,0,0,.1));background-image:-moz-linear-gradient(top,rgba(0,0,0,.05) 0,rgba(0,0,0,.1));background-image:-o-linear-gradient(transparent,rgba(0,0,0,.05) 40%,rgba(0,0,0,.1));background-image:linear-gradient(transparent,rgba(0,0,0,.05) 40%,rgba(0,0,0,.1))}.pure-button:focus{outline:0}.pure-button-active,.pure-button:active{box-shadow:0 0 0 1px rgba(0,0,0,.15) inset,0 0 6px rgba(0,0,0,.2) inset;border-color:#000\9}.pure-button[disabled],.pure-button-disabled,.pure-button-disabled:hover,.pure-button-disabled:focus,.pure-button-disabled:active{border:0;background-image:none;filter:progid:DXImageTransform.Microsoft.gradient(enabled=false);filter:alpha(opacity=40);-khtml-opacity:.4;-moz-opacity:.4;opacity:.4;cursor:not-allowed;box-shadow:none}.pure-button-hidden{display:none}.pure-button::-moz-focus-inner{padding:0;border:0}.pure-button-primary,.pure-button-selected,a.pure-button-primary,a.pure-button-selected{background-color:#0078e7;color:#fff}.pure-form input[type=text],.pure-form input[type=password],.pure-form input[type=email],.pure-form input[type=url],.pure-form input[type=date],.pure-form input[type=month],.pure-form input[type=time],.pure-form input[type=datetime],.pure-form input[type=datetime-local],.pure-form input[type=week],.pure-form input[type=number],.pure-form input[type=search],.pure-form input[type=tel],.pure-form input[type=color],.pure-form select,.pure-form textarea{padding:.5em .6em;display:inline-block;border:1px solid #ccc;box-shadow:inset 0 1px 3px #ddd;border-radius:4px;vertical-align:middle;-webkit-box-sizing:border-box;-moz-box-sizing:border-box;box-sizing:border-box}.pure-form input:not([type]){padding:.5em .6em;display:inline-block;border:1px solid #ccc;box-shadow:inset 0 1px 3px #ddd;border-radius:4px;-webkit-box-sizing:border-box;-moz-box-sizing:border-box;box-sizing:border-box}.pure-form input[type=color]{padding:.2em .5em}.pure-form input[type=text]:focus,.pure-form input[type=password]:focus,.pure-form input[type=email]:focus,.pure-form input[type=url]:focus,.pure-form input[type=date]:focus,.pure-form input[type=month]:focus,.pure-form input[type=time]:focus,.pure-form input[type=datetime]:focus,.pure-form input[type=datetime-local]:focus,.pure-form input[type=week]:focus,.pure-form input[type=number]:focus,.pure-form input[type=search]:focus,.pure-form input[type=tel]:focus,.pure-form input[type=color]:focus,.pure-form select:focus,.pure-form textarea:focus{outline:0;border-color:#129FEA}.pure-form input:not([type]):focus{outline:0;border-color:#129FEA}.pure-form input[type=file]:focus,.pure-form input[type=radio]:focus,.pure-form input[type=checkbox]:focus{outline:thin solid #129FEA;outline:1px auto #129FEA}.pure-form .pure-checkbox,.pure-form .pure-radio{margin:.5em 0;display:block}.pure-form input[type=text][disabled],.pure-form input[type=password][disabled],.pure-form input[type=email][disabled],.pure-form input[type=url][disabled],.pure-form input[type=date][disabled],.pure-form input[type=month][disabled],.pure-form input[type=time][disabled],.pure-form input[type=datetime][disabled],.pure-form input[type=datetime-local][disabled],.pure-form input[type=week][disabled],.pure-form input[type=number][disabled],.pure-form input[type=search][disabled],.pure-form input[type=tel][disabled],.pure-form input[type=color][disabled],.pure-form select[disabled],.pure-form textarea[disabled]{cursor:not-allowed;background-color:#eaeded;color:#cad2d3}.pure-form input:not([type])[disabled]{cursor:not-allowed;background-color:#eaeded;color:#cad2d3}.pure-form input[readonly],.pure-form select[readonly],.pure-form textarea[readonly]{background-color:#eee;color:#777;border-color:#ccc}.pure-form input:focus:invalid,.pure-form textarea:focus:invalid,.pure-form select:focus:invalid{color:#b94a48;border-color:#e9322d}.pure-form input[type=file]:focus:invalid:focus,.pure-form input[type=radio]:focus:invalid:focus,.pure-form input[type=checkbox]:focus:invalid:focus{outline-color:#e9322d}.pure-form select{height:2.25em;border:1px solid #ccc;background-color:#fff}.pure-form select[multiple]{height:auto}.pure-form label{margin:.5em 0 .2em}.pure-form fieldset{margin:0;padding:.35em 0 .75em;border:0}.pure-form legend{display:block;width:100%;padding:.3em 0;margin-bottom:.3em;color:#333;border-bottom:1px solid #e5e5e5}.pure-form-stacked input[type=text],.pure-form-stacked input[type=password],.pure-form-stacked input[type=email],.pure-form-stacked input[type=url],.pure-form-stacked input[type=date],.pure-form-stacked input[type=month],.pure-form-stacked input[type=time],.pure-form-stacked input[type=datetime],.pure-form-stacked input[type=datetime-local],.pure-form-stacked input[type=week],.pure-form-stacked input[type=number],.pure-form-stacked input[type=search],.pure-form-stacked input[type=tel],.pure-form-stacked input[type=color],.pure-form-stacked input[type=file],.pure-form-stacked select,.pure-form-stacked label,.pure-form-stacked textarea{display:block;margin:.25em 0}.pure-form-stacked input:not([type]){display:block;margin:.25em 0}.pure-form-aligned input,.pure-form-aligned textarea,.pure-form-aligned select,.pure-form-aligned .pure-help-inline,.pure-form-message-inline{display:inline-block;*display:inline;*zoom:1;vertical-align:middle}.pure-form-aligned textarea{vertical-align:top}.pure-form-aligned .pure-control-group{margin-bottom:.5em}.pure-form-aligned .pure-control-group label{text-align:right;display:inline-block;vertical-align:middle;width:10em;margin:0 1em 0 0}.pure-form-aligned .pure-controls{margin:1.5em 0 0 11em}.pure-form input.pure-input-rounded,.pure-form .pure-input-rounded{border-radius:2em;padding:.5em 1em}.pure-form .pure-group fieldset{margin-bottom:10px}.pure-form .pure-group input,.pure-form .pure-group textarea{display:block;padding:10px;margin:0 0 -1px;border-radius:0;position:relative;top:-1px}.pure-form .pure-group input:focus,.pure-form .pure-group textarea:focus{z-index:3}.pure-form .pure-group input:first-child,.pure-form .pure-group textarea:first-child{top:1px;border-radius:4px 4px 0 0;margin:0}.pure-form .pure-group input:first-child:last-child,.pure-form .pure-group textarea:first-child:last-child{top:1px;border-radius:4px;margin:0}.pure-form .pure-group input:last-child,.pure-form .pure-group textarea:last-child{top:-2px;border-radius:0 0 4px 4px;margin:0}.pure-form .pure-group button{margin:.35em 0}.pure-form .pure-input-1{width:100%}.pure-form .pure-input-2-3{width:66%}.pure-form .pure-input-1-2{width:50%}.pure-form .pure-input-1-3{width:33%}.pure-form .pure-input-1-4{width:25%}.pure-form .pure-help-inline,.pure-form-message-inline{display:inline-block;padding-left:.3em;color:#666;vertical-align:middle;font-size:.875em}.pure-form-message{display:block;color:#666;font-size:.875em}@media only screen and (max-width :480px){.pure-form button[type=submit]{margin:.7em 0 0}.pure-form input:not([type]),.pure-form input[type=text],.pure-form input[type=password],.pure-form input[type=email],.pure-form input[type=url],.pure-form input[type=date],.pure-form input[type=month],.pure-form input[type=time],.pure-form input[type=datetime],.pure-form input[type=datetime-local],.pure-form input[type=week],.pure-form input[type=number],.pure-form input[type=search],.pure-form input[type=tel],.pure-form input[type=color],.pure-form label{margin-bottom:.3em;display:block}.pure-group input:not([type]),.pure-group input[type=text],.pure-group input[type=password],.pure-group input[type=email],.pure-group input[type=url],.pure-group input[type=date],.pure-group input[type=month],.pure-group input[type=time],.pure-group input[type=datetime],.pure-group input[type=datetime-local],.pure-group input[type=week],.pure-group input[type=number],.pure-group input[type=search],.pure-group input[type=tel],.pure-group input[type=color]{margin-bottom:0}.pure-form-aligned .pure-control-group label{margin-bottom:.3em;text-align:left;display:block;width:100%}.pure-form-aligned .pure-controls{margin:1.5em 0 0}.pure-form .pure-help-inline,.pure-form-message-inline,.pure-form-message{display:block;font-size:.75em;padding:.2em 0 .8em}}.pure-menu{-webkit-box-sizing:border-box;-moz-box-sizing:border-box;box-sizing:border-box}.pure-menu-fixed{position:fixed;left:0;top:0;z-index:3}.pure-menu-list,.pure-menu-item{position:relative}.pure-menu-list{list-style:none;margin:0;padding:0}.pure-menu-item{padding:0;margin:0;height:100%}.pure-menu-link,.pure-menu-heading{display:block;text-decoration:none;white-space:nowrap}.pure-menu-horizontal{width:100%;white-space:nowrap}.pure-menu-horizontal .pure-menu-list{display:inline-block}.pure-menu-horizontal .pure-menu-item,.pure-menu-horizontal .pure-menu-heading,.pure-menu-horizontal .pure-menu-separator{display:inline-block;*display:inline;zoom:1;vertical-align:middle}.pure-menu-item .pure-menu-item{display:block}.pure-menu-children{display:none;position:absolute;left:100%;top:0;margin:0;padding:0;z-index:3}.pure-menu-horizontal .pure-menu-children{left:0;top:auto;width:inherit}.pure-menu-allow-hover:hover>.pure-menu-children,.pure-menu-active>.pure-menu-children{display:block;position:absolute}.pure-menu-has-children>.pure-menu-link:after{padding-left:.5em;content:"\25B8";font-size:small}.pure-menu-horizontal .pure-menu-has-children>.pure-menu-link:after{content:"\25BE"}.pure-menu-scrollable{overflow-y:scroll;overflow-x:hidden}.pure-menu-scrollable .pure-menu-list{display:block}.pure-menu-horizontal.pure-menu-scrollable .pure-menu-list{display:inline-block}.pure-menu-horizontal.pure-menu-scrollable{white-space:nowrap;overflow-y:hidden;overflow-x:auto;-ms-overflow-style:none;-webkit-overflow-scrolling:touch;padding:.5em 0}.pure-menu-horizontal.pure-menu-scrollable::-webkit-scrollbar{display:none}.pure-menu-separator{background-color:#ccc;height:1px;margin:.3em 0}.pure-menu-horizontal .pure-menu-separator{width:1px;height:1.3em;margin:0 .3em}.pure-menu-heading{text-transform:uppercase;color:#565d64}.pure-menu-link{color:#777}.pure-menu-children{background-color:#fff}.pure-menu-link,.pure-menu-disabled,.pure-menu-heading{padding:.5em 1em}.pure-menu-disabled{opacity:.5}.pure-menu-disabled .pure-menu-link:hover{background-color:transparent}.pure-menu-active>.pure-menu-link,.pure-menu-link:hover,.pure-menu-link:focus{background-color:#eee}.pure-menu-selected .pure-menu-link,.pure-menu-selected .pure-menu-link:visited{color:#000}.pure-table{border-collapse:collapse;border-spacing:0;empty-cells:show;border:1px solid #cbcbcb}.pure-table caption{color:#000;font:italic 85%/1 arial,sans-serif;padding:1em 0;text-align:center}.pure-table td,.pure-table th{border-left:1px solid #cbcbcb;border-width:0 0 0 1px;font-size:inherit;margin:0;overflow:visible;padding:.5em 1em}.pure-table td:first-child,.pure-table th:first-child{border-left-width:0}.pure-table thead{background-color:#e0e0e0;color:#000;text-align:left;vertical-align:bottom}.pure-table td{background-color:transparent}.pure-table-odd td{background-color:#f2f2f2}.pure-table-striped tr:nth-child(2n-1) td{background-color:#f2f2f2}.pure-table-bordered td{border-bottom:1px solid #cbcbcb}.pure-table-bordered tbody>tr:last-child>td{border-bottom-width:0}.pure-table-horizontal td,.pure-table-horizontal th{border-width:0 0 1px;border-bottom:1px solid #cbcbcb}.pure-table-horizontal tbody>tr:last-child>td{border-bottom-width:0}
`)

var Scooch_style_min_css = []byte(`
.m-scooch-controls{padding-top:10px;text-align:center}.m-scooch-controls a{padding:5px;-webkit-user-select:none;-moz-user-select:-moz-none;user-select:none;-webkit-user-drag:none;-moz-user-drag:-moz-none;user-drag:none}.m-scooch-bulleted a{line-height:0;text-decoration:none;text-indent:-999px;overflow:hidden;display:inline-block;padding:6px;width:0;height:0;margin:0 3px;color:#333;background-color:rgba(255,255,255,.3);-webkit-transition:background-color .1s ease-in;-moz-transition:background-color .1s ease-in;-o-transition:background-color .1s ease-in;transition:background-color .1s ease-in;-webkit-box-shadow:inset rgba(0,0,0,.25) 0 1px 2px;-moz-box-shadow:inset rgba(0,0,0,.25) 0 1px 2px;box-shadow:inset rgba(0,0,0,.25) 0 1px 2px;-webkit-border-radius:6px;-moz-border-radius:6px;border-radius:6px}.m-scooch-bulleted a:hover,.m-scooch-bulleted a:focus{text-decoration:none;background-color:rgba(255,255,255,.6)}.m-scooch-bulleted a.m-active{background-color:rgba(255,255,255,1);-webkit-box-shadow:rgba(0,0,0,.25) 0 1px 2px;-moz-box-shadow:rgba(0,0,0,.25) 0 1px 2px;box-shadow:rgba(0,0,0,.25) 0 1px 2px}.m-scooch-pagination{padding-top:10px}.m-scooch-pagination a{text-decoration:none;display:inline-block;padding:3px 10px;margin:1px 0;color:#333;background-color:rgba(255,255,255,.3);-webkit-transition:background-color .1s ease-in;-moz-transition:background-color .1s ease-in;-o-transition:background-color .1s ease-in;transition:background-color .1s ease-in;-webkit-border-radius:2px;-moz-border-radius:2px;border-radius:2px}.m-scooch-pagination a:hover,.m-scooch-pagination a:focus{text-decoration:none;background-color:rgba(255,255,255,.6)}.m-scooch-pagination a.m-active{background-color:rgba(255,255,255,1)}.m-scooch-hud{padding-top:0}.m-scooch-hud a{z-index:2;opacity:0;display:block;position:absolute;top:50%;width:50px;height:50px;margin:-25px 0 0 0;padding:0;text-decoration:none;text-indent:-999px;overflow:hidden;color:rgba(255,255,255,.8);background:rgba(0,0,0,.8);-webkit-transition:opacity .1s ease-in;-moz-transition:opacity .1s ease-in;-o-transition:opacity .1s ease-in;transition:opacity .1s ease-in;-webkit-border-radius:25px;-moz-border-radius:25px;border-radius:25px}.m-scooch:hover .m-scooch-hud a{opacity:.3}.m-scooch .m-scooch-hud a:hover,.m-scooch .m-scooch-hud a:focus{opacity:1}.m-scooch-hud a:after{color:rgba(255,255,255,.85);content:"\25c0";font-size:25px;font-weight:700;text-indent:0;text-align:center;display:block;position:absolute;top:10px;left:0;width:47px;height:50px;z-index:9}.m-scooch-hud .m-scooch-prev{left:10px}.m-scooch-hud .m-scooch-next{right:10px}.m-scooch-hud .m-scooch-next:after{left:auto;right:0;content:"\25b6"}.m-caption{margin:0;padding:10px;height:auto;text-align:center}.m-scaled .m-item{opacity:.7;-webkit-backface-visibility:hidden;-webkit-transform:scale(0.75);-moz-transform:scale(0.75);-ms-transform:scale(0.75);-o-transform:scale(0.75);transform:scale(0.75);-webkit-transition:-webkit-transform cubic-bezier(0.33,.66,.66,1) .25s,opacity ease-out .25s;-moz-transition-timing-function:-moz-transform cubic-bezier(0.33,.66,.66,1) .25s,opacity ease-out .25s;-o-transition-timing-function:-o-transform cubic-bezier(0.33,.66,.66,1) .25s,opacity ease-out .25s;transition-timing-function:transform cubic-bezier(0.33,.66,.66,1) .25s,opacity ease-out .25s}.m-scaled .m-active{opacity:1;-webkit-transform:scale(1);-moz-transform:scale(1);-ms-transform:scale(1);-o-transform:scale(1);transform:scale(1)}.m-fluid .m-item{margin-right:20px}.m-scooch-photos{margin:0 -10px;padding:0 10px}.m-scooch-photos .m-item>img{margin:0;padding:0;max-width:none;width:100%;height:auto;-webkit-box-shadow:rgba(0,0,0,.5) 0 5px 10px;-moz-box-shadow:rgba(0,0,0,.5) 0 5px 10px;-o-box-shadow:rgba(0,0,0,.5) 0 5px 10px;-ms-box-shadow:rgba(0,0,0,.5) 0 5px 10px;box-shadow:rgba(0,0,0,.5) 0 5px 10px}.m-scooch-photos .m-caption{background:rgba(0,0,0,.7);bottom:0;position:absolute;z-index:9;width:100%;-webkit-box-sizing:border-box;-moz-box-sizing:border-box;box-sizing:border-box}.m-card-dark,.m-card-light{padding:20px;-webkit-border-radius:6px;-moz-border-radius:6px;border-radius:6px;-webkit-box-shadow:rgba(0,0,0,.5) 0 5px 10px;-moz-box-shadow:rgba(0,0,0,.5) 0 5px 10px;-o-box-shadow:rgba(0,0,0,.5) 0 5px 10px;-ms-box-shadow:rgba(0,0,0,.5) 0 5px 10px;box-shadow:rgba(0,0,0,.5) 0 5px 10px}.m-card-dark{background:rgba(0,0,0,.5);color:#FFF}.m-card-light{background:rgba(255,255,255,.9);color:#000}.m-card-dark .m-caption,.m-card-light .m-caption{margin:0;padding:10px 0 0}.m-fade-out{-webkit-mask-image:-webkit-linear-gradient(left,rgba(0,0,0,0) 0,rgba(0,0,0,1) 5%,rgba(0,0,0,1) 95%,rgba(0,0,0,0) 100%)}
`)

var Scooch_min_css = []byte(`
.m-scooch{position:relative;overflow:hidden;-webkit-font-smoothing:antialiased}.m-scooch.m-left{text-align:left}.m-scooch.m-center{text-align:center}.m-scooch.m-fluid>.m-scooch-inner>*{width:100%}.m-scooch.m-fluid.m-center>.m-scooch-inner>:first-child{margin-left:0}.m-scooch.m-fluid-2>.m-scooch-inner>*{width:50%}.m-scooch.m-fluid-2.m-center>.m-scooch-inner>:first-child{margin-left:25%}.m-scooch.m-fluid-3>.m-scooch-inner>*{width:33.333%}.m-scooch.m-fluid-3.m-center>.m-scooch-inner>:first-child{margin-left:33.333%}.m-scooch.m-fluid-4>.m-scooch-inner>*{width:25%}.m-scooch.m-fluid-4.m-center>.m-scooch-inner>:first-child{margin-left:37.5%}.m-scooch.m-fluid-5>.m-scooch-inner>*{width:20%}.m-scooch.m-fluid-5.m-center>.m-scooch-inner>:first-child{margin-left:40%}.m-scooch.m-fluid-6>.m-scooch-inner>*{width:16.667%}.m-scooch.m-fluid-6.m-center>.m-scooch-inner>:first-child{margin-left:41.667%}.m-scooch img{-ms-interpolation-mode:bicubic}.m-scooch .m-item{-webkit-transform:translate(0);transform:translate(0)}.m-scooch-inner{position:relative;white-space:nowrap;text-align:left;font-size:0;-webkit-transition-property:-webkit-transform;-moz-transition-property:-moz-transform;-ms-transition-property:-ms-transform;-o-transition-property:-o-transform;transition-property:transform;-webkit-transition-timing-function:cubic-bezier(0.33,.66,.66,1);-moz-transition-timing-function:cubic-bezier(0.33,.66,.66,1);-ms-transition-timing-function:cubic-bezier(0.33,.66,.66,1);-o-transition-timing-function:cubic-bezier(0.33,.66,.66,1);transition-timing-function:cubic-bezier(0.33,.66,.66,1);-webkit-transition-duration:.5s;-moz-transition-duration:.5s;-ms-transition-duration:.5s;-o-transition-duration:.5s;transition-duration:.5s}.m-scooch-inner>*{display:inline-block;vertical-align:top;white-space:normal;font-size:16px}.m-fluid>.m-scooch-inner>*{box-sizing:border-box;-ms-box-sizing:border-box;-moz-box-sizing:border-box;-o-box-sizing:border-box;-webkit-box-sizing:border-box}.m-center:not(.m-fluid)>.m-scooch-inner{display:inline-block;margin-right:-20000px!important;margin-left:0!important}.m-center:not(.m-fluid)>.m-scooch-inner>*{position:relative;left:-20000px}.m-center:not(.m-fluid)>.m-scooch-inner>:first-child{float:left;margin-right:20000px;left:0}.m-center:not(.m-fluid)>.m-scooch-inner>:first-child:last-child{margin-right:0}.m-center:not(.m-fluid)>.m-scooch-inner>:last-child{margin-right:-30000px}
`)

var Scooch_min_js = []byte(`
/*! scooch 0.5.0 2015-01-08 */
(function(t){if("function"==typeof define&&define.amd)define(["$"],t);else{var e=window.Mobify&&window.Mobify.$||window.Zepto||window.jQuery;t(e)}})(function(t){var e=function(t){var e={},i=navigator.userAgent,n=t.support=t.support||{};t.extend(t.support,{touch:"ontouchend"in document}),e.events=n.touch?{down:"touchstart",move:"touchmove",up:"touchend"}:{down:"mousedown",move:"mousemove",up:"mouseup"},e.getCursorPosition=n.touch?function(t){return t=t.originalEvent||t,{x:t.touches[0].clientX,y:t.touches[0].clientY}}:function(t){return{x:t.clientX,y:t.clientY}},e.getProperty=function(t){for(var e=["Webkit","Moz","O","ms",""],i=document.createElement("div").style,n=0;e.length>n;++n)if(void 0!==i[e[n]+t])return e[n]+t},t.extend(n,{transform:!!e.getProperty("Transform"),transform3d:!(!(window.WebKitCSSMatrix&&"m11"in new WebKitCSSMatrix)||/android\s+[1-2]/i.test(i))});var s=e.getProperty("Transform");e.translateX=n.transform3d?function(t,e){"number"==typeof e&&(e+="px"),t.style[s]="translate3d("+e+",0,0)"}:n.transform?function(t,e){"number"==typeof e&&(e+="px"),t.style[s]="translate("+e+",0)"}:function(t,e){"number"==typeof e&&(e+="px"),t.style.left=e};var o=(e.getProperty("Transition"),e.getProperty("TransitionDuration"));return e.setTransitions=function(t,e){t.style[o]=e?"":"0s"},e.requestAnimationFrame=function(){var t=window.requestAnimationFrame||window.webkitRequestAnimationFrame||window.mozRequestAnimationFrame||window.oRequestAnimationFrame||window.msRequestAnimationFrame||function(t){window.setTimeout(t,1e3/60)},e=function(){t.apply(window,arguments)};return e}(),e}(t),i=function(t,e){var i={dragRadius:10,moveRadius:20,animate:!0,autoHideArrows:!1,rightToLeft:!1,classPrefix:"m-",classNames:{outer:"scooch",inner:"scooch-inner",item:"item",center:"center",touch:"has-touch",dragging:"dragging",active:"active",inactive:"inactive",fluid:"fluid"}},n=t.support,s=function(t,e){this.setOptions(e),this.initElements(t),this.initOffsets(),this.initAnimation(),this.bind(),this._updateCallbacks=[]};return s.defaults=i,s.prototype.setOptions=function(e){var n=this.options||t.extend({},i,e);n.classNames=t.extend({},n.classNames,e.classNames||{}),this.options=n},s.prototype.initElements=function(e){this._index=1,this.element=e,this.$element=t(e),this.$inner=this.$element.find("."+this._getClass("inner")),this.$items=this.$inner.children(),this.$start=this.$items.eq(0),this.$sec=this.$items.eq(1),this.$current=this.$items.eq(this._index-1),this._length=this.$items.length,this._alignment=this.$element.hasClass(this._getClass("center"))?.5:0,this._isFluid=this.$element.hasClass(this._getClass("fluid"))},s.prototype.initOffsets=function(){this._offsetDrag=0},s.prototype.initAnimation=function(){this.animating=!1,this.dragging=!1,this._needsUpdate=!1,this._enableAnimation()},s.prototype._getClass=function(t){return this.options.classPrefix+this.options.classNames[t]},s.prototype._enableAnimation=function(){this.animating||(e.setTransitions(this.$inner[0],!0),this.$inner.removeClass(this._getClass("dragging")),this.animating=!0)},s.prototype._disableAnimation=function(){this.animating&&(e.setTransitions(this.$inner[0],!1),this.$inner.addClass(this._getClass("dragging")),this.animating=!1)},s.prototype.refresh=function(){this.$items=this.$inner.children("."+this._getClass("item")),this.$start=this.$items.eq(0),this.$sec=this.$items.eq(1),this._length=this.$items.length,this.update()},s.prototype.update=function(t){if(t!==void 0&&this._updateCallbacks.push(t),!this._needsUpdate){this._needsUpdate=!0;var i=this;e.requestAnimationFrame(function(){i._update(),setTimeout(function(){for(var t=0,e=i._updateCallbacks.length;e>t;t++)i._updateCallbacks[t].call(i);i._updateCallbacks=[]},10)})}},s.prototype._update=function(){if(this._needsUpdate){var t=this.$current,i=this.$start,n=t.prop("offsetLeft")+t.prop("clientWidth")*this._alignment,s=i.prop("offsetLeft")+i.prop("clientWidth")*this._alignment,o=Math.round(-(n-s)+this._offsetDrag);e.translateX(this.$inner[0],o),this._needsUpdate=!1}},s.prototype.bind=function(){function i(t){n.touch||t.preventDefault(),c=!0,m=!1,r=e.getCursorPosition(t),h=0,d=0,l=!1,p._disableAnimation(),$=1==p._index,y=p._index==p._length}function s(t){if(c&&!m){var i=e.getCursorPosition(t),n=p.$element.width();h=r.x-i.x,d=r.y-i.y,l||u(h)>u(d)&&u(h)>f?(l=!0,t.preventDefault(),$&&0>h?h=h*-n/(h-n):y&&h>0&&(h=h*n/(h+n)),p._offsetDrag=-h,p.update()):u(d)>u(h)&&u(d)>f&&(m=!0)}}function o(){c&&(c=!1,p._enableAnimation(),!m&&u(h)>_.moveRadius?_.rightToLeft?0>h?p.next():p.prev():h>0?p.next():p.prev():(p._offsetDrag=0,p.update()))}function a(t){l&&t.preventDefault()}var r,h,d,l,u=Math.abs,c=!1,m=!1,f=this.options.dragRadius,p=this,g=this.$element,v=this.$inner,_=this.options,$=!1,y=!1,w=t(window).width();v.on(e.events.down+".scooch",i).on(e.events.move+".scooch",s).on(e.events.up+".scooch",o).on("click.scooch",a).on("mouseout.scooch",o),g.on("click","[data-m-slide]",function(e){e.preventDefault();var i=t(this).attr("data-m-slide"),n=parseInt(i,10);isNaN(n)?p[i]():p.move(n)}),g.on("afterSlide",function(t,e,i){p.$items.eq(e-1).removeClass(p._getClass("active")),p.$items.eq(i-1).addClass(p._getClass("active")),p.$element.find("[data-m-slide='"+e+"']").removeClass(p._getClass("active")),p.$element.find("[data-m-slide='"+i+"']").addClass(p._getClass("active")),_.autoHideArrows&&(p.$element.find("[data-m-slide=prev]").removeClass(p._getClass("inactive")),p.$element.find("[data-m-slide=next]").removeClass(p._getClass("inactive")),1===i&&p.$element.find("[data-m-slide=prev]").addClass(p._getClass("inactive")),i===p._length&&p.$element.find("[data-m-slide=next]").addClass(p._getClass("inactive")))}),t(window).on("resize orientationchange",function(){w!=t(window).width()&&(p._disableAnimation(),w=t(window).width(),p.update())}),g.trigger("beforeSlide",[1,1]),g.trigger("afterSlide",[1,1]),p.update()},s.prototype.unbind=function(){this.$inner.off()},s.prototype.destroy=function(){this.unbind(),this.$element.trigger("destroy"),this.$element.remove(),this.$element=null,this.$inner=null,this.$start=null,this.$current=null},s.prototype.move=function(e,i){var n=this.$element,s=(this.$inner,this.$items),o=(this.$start,this.$current),a=this._length,r=this._index;i=t.extend({},this.options,i),1>e?e=1:e>this._length&&(e=a),e==this._index,i.animate?this._enableAnimation():this._disableAnimation(),n.trigger("beforeSlide",[r,e]),this.$current=o=s.eq(e-1),this._offsetDrag=0,this._index=e,i.animate?this.update():this.update(function(){this._enableAnimation()}),n.trigger("afterSlide",[r,e])},s.prototype.next=function(){this.move(this._index+1)},s.prototype.prev=function(){this.move(this._index-1)},s}(t,e);t.fn.scooch=function(e,n){var s=t.extend({},t.fn.scooch.defaults);return"object"==typeof e&&(t.extend(s,e,!0),n=null,e=null),n=Array.prototype.slice.apply(arguments),this.each(function(){var o=(t(this),this._scooch);o||(o=new i(this,s)),e&&(o[e].apply(o,n.slice(1)),"destroy"===e&&(o=null)),this._scooch=o}),this},t.fn.scooch.defaults={}});
`)

var Zepto_min_js = []byte(`
/* Zepto v1.1.6 - zepto event ajax form ie - zeptojs.com/license */
var Zepto=function(){function L(t){return null==t?String(t):j[S.call(t)]||"object"}function Z(t){return"function"==L(t)}function _(t){return null!=t&&t==t.window}function $(t){return null!=t&&t.nodeType==t.DOCUMENT_NODE}function D(t){return"object"==L(t)}function M(t){return D(t)&&!_(t)&&Object.getPrototypeOf(t)==Object.prototype}function R(t){return"number"==typeof t.length}function k(t){return s.call(t,function(t){return null!=t})}function z(t){return t.length>0?n.fn.concat.apply([],t):t}function F(t){return t.replace(/::/g,"/").replace(/([A-Z]+)([A-Z][a-z])/g,"$1_$2").replace(/([a-z\d])([A-Z])/g,"$1_$2").replace(/_/g,"-").toLowerCase()}function q(t){return t in f?f[t]:f[t]=new RegExp("(^|\\s)"+t+"(\\s|$)")}function H(t,e){return"number"!=typeof e||c[F(t)]?e:e+"px"}function I(t){var e,n;return u[t]||(e=a.createElement(t),a.body.appendChild(e),n=getComputedStyle(e,"").getPropertyValue("display"),e.parentNode.removeChild(e),"none"==n&&(n="block"),u[t]=n),u[t]}function V(t){return"children"in t?o.call(t.children):n.map(t.childNodes,function(t){return 1==t.nodeType?t:void 0})}function B(n,i,r){for(e in i)r&&(M(i[e])||A(i[e]))?(M(i[e])&&!M(n[e])&&(n[e]={}),A(i[e])&&!A(n[e])&&(n[e]=[]),B(n[e],i[e],r)):i[e]!==t&&(n[e]=i[e])}function U(t,e){return null==e?n(t):n(t).filter(e)}function J(t,e,n,i){return Z(e)?e.call(t,n,i):e}function X(t,e,n){null==n?t.removeAttribute(e):t.setAttribute(e,n)}function W(e,n){var i=e.className||"",r=i&&i.baseVal!==t;return n===t?r?i.baseVal:i:void(r?i.baseVal=n:e.className=n)}function Y(t){try{return t?"true"==t||("false"==t?!1:"null"==t?null:+t+""==t?+t:/^[\[\{]/.test(t)?n.parseJSON(t):t):t}catch(e){return t}}function G(t,e){e(t);for(var n=0,i=t.childNodes.length;i>n;n++)G(t.childNodes[n],e)}var t,e,n,i,C,N,r=[],o=r.slice,s=r.filter,a=window.document,u={},f={},c={"column-count":1,columns:1,"font-weight":1,"line-height":1,opacity:1,"z-index":1,zoom:1},l=/^\s*<(\w+|!)[^>]*>/,h=/^<(\w+)\s*\/?>(?:<\/\1>|)$/,p=/<(?!area|br|col|embed|hr|img|input|link|meta|param)(([\w:]+)[^>]*)\/>/gi,d=/^(?:body|html)$/i,m=/([A-Z])/g,g=["val","css","html","text","data","width","height","offset"],v=["after","prepend","before","append"],y=a.createElement("table"),x=a.createElement("tr"),b={tr:a.createElement("tbody"),tbody:y,thead:y,tfoot:y,td:x,th:x,"*":a.createElement("div")},w=/complete|loaded|interactive/,E=/^[\w-]*$/,j={},S=j.toString,T={},O=a.createElement("div"),P={tabindex:"tabIndex",readonly:"readOnly","for":"htmlFor","class":"className",maxlength:"maxLength",cellspacing:"cellSpacing",cellpadding:"cellPadding",rowspan:"rowSpan",colspan:"colSpan",usemap:"useMap",frameborder:"frameBorder",contenteditable:"contentEditable"},A=Array.isArray||function(t){return t instanceof Array};return T.matches=function(t,e){if(!e||!t||1!==t.nodeType)return!1;var n=t.webkitMatchesSelector||t.mozMatchesSelector||t.oMatchesSelector||t.matchesSelector;if(n)return n.call(t,e);var i,r=t.parentNode,o=!r;return o&&(r=O).appendChild(t),i=~T.qsa(r,e).indexOf(t),o&&O.removeChild(t),i},C=function(t){return t.replace(/-+(.)?/g,function(t,e){return e?e.toUpperCase():""})},N=function(t){return s.call(t,function(e,n){return t.indexOf(e)==n})},T.fragment=function(e,i,r){var s,u,f;return h.test(e)&&(s=n(a.createElement(RegExp.$1))),s||(e.replace&&(e=e.replace(p,"<$1></$2>")),i===t&&(i=l.test(e)&&RegExp.$1),i in b||(i="*"),f=b[i],f.innerHTML=""+e,s=n.each(o.call(f.childNodes),function(){f.removeChild(this)})),M(r)&&(u=n(s),n.each(r,function(t,e){g.indexOf(t)>-1?u[t](e):u.attr(t,e)})),s},T.Z=function(t,e){return t=t||[],t.__proto__=n.fn,t.selector=e||"",t},T.isZ=function(t){return t instanceof T.Z},T.init=function(e,i){var r;if(!e)return T.Z();if("string"==typeof e)if(e=e.trim(),"<"==e[0]&&l.test(e))r=T.fragment(e,RegExp.$1,i),e=null;else{if(i!==t)return n(i).find(e);r=T.qsa(a,e)}else{if(Z(e))return n(a).ready(e);if(T.isZ(e))return e;if(A(e))r=k(e);else if(D(e))r=[e],e=null;else if(l.test(e))r=T.fragment(e.trim(),RegExp.$1,i),e=null;else{if(i!==t)return n(i).find(e);r=T.qsa(a,e)}}return T.Z(r,e)},n=function(t,e){return T.init(t,e)},n.extend=function(t){var e,n=o.call(arguments,1);return"boolean"==typeof t&&(e=t,t=n.shift()),n.forEach(function(n){B(t,n,e)}),t},T.qsa=function(t,e){var n,i="#"==e[0],r=!i&&"."==e[0],s=i||r?e.slice(1):e,a=E.test(s);return $(t)&&a&&i?(n=t.getElementById(s))?[n]:[]:1!==t.nodeType&&9!==t.nodeType?[]:o.call(a&&!i?r?t.getElementsByClassName(s):t.getElementsByTagName(e):t.querySelectorAll(e))},n.contains=a.documentElement.contains?function(t,e){return t!==e&&t.contains(e)}:function(t,e){for(;e&&(e=e.parentNode);)if(e===t)return!0;return!1},n.type=L,n.isFunction=Z,n.isWindow=_,n.isArray=A,n.isPlainObject=M,n.isEmptyObject=function(t){var e;for(e in t)return!1;return!0},n.inArray=function(t,e,n){return r.indexOf.call(e,t,n)},n.camelCase=C,n.trim=function(t){return null==t?"":String.prototype.trim.call(t)},n.uuid=0,n.support={},n.expr={},n.map=function(t,e){var n,r,o,i=[];if(R(t))for(r=0;r<t.length;r++)n=e(t[r],r),null!=n&&i.push(n);else for(o in t)n=e(t[o],o),null!=n&&i.push(n);return z(i)},n.each=function(t,e){var n,i;if(R(t)){for(n=0;n<t.length;n++)if(e.call(t[n],n,t[n])===!1)return t}else for(i in t)if(e.call(t[i],i,t[i])===!1)return t;return t},n.grep=function(t,e){return s.call(t,e)},window.JSON&&(n.parseJSON=JSON.parse),n.each("Boolean Number String Function Array Date RegExp Object Error".split(" "),function(t,e){j["[object "+e+"]"]=e.toLowerCase()}),n.fn={forEach:r.forEach,reduce:r.reduce,push:r.push,sort:r.sort,indexOf:r.indexOf,concat:r.concat,map:function(t){return n(n.map(this,function(e,n){return t.call(e,n,e)}))},slice:function(){return n(o.apply(this,arguments))},ready:function(t){return w.test(a.readyState)&&a.body?t(n):a.addEventListener("DOMContentLoaded",function(){t(n)},!1),this},get:function(e){return e===t?o.call(this):this[e>=0?e:e+this.length]},toArray:function(){return this.get()},size:function(){return this.length},remove:function(){return this.each(function(){null!=this.parentNode&&this.parentNode.removeChild(this)})},each:function(t){return r.every.call(this,function(e,n){return t.call(e,n,e)!==!1}),this},filter:function(t){return Z(t)?this.not(this.not(t)):n(s.call(this,function(e){return T.matches(e,t)}))},add:function(t,e){return n(N(this.concat(n(t,e))))},is:function(t){return this.length>0&&T.matches(this[0],t)},not:function(e){var i=[];if(Z(e)&&e.call!==t)this.each(function(t){e.call(this,t)||i.push(this)});else{var r="string"==typeof e?this.filter(e):R(e)&&Z(e.item)?o.call(e):n(e);this.forEach(function(t){r.indexOf(t)<0&&i.push(t)})}return n(i)},has:function(t){return this.filter(function(){return D(t)?n.contains(this,t):n(this).find(t).size()})},eq:function(t){return-1===t?this.slice(t):this.slice(t,+t+1)},first:function(){var t=this[0];return t&&!D(t)?t:n(t)},last:function(){var t=this[this.length-1];return t&&!D(t)?t:n(t)},find:function(t){var e,i=this;return e=t?"object"==typeof t?n(t).filter(function(){var t=this;return r.some.call(i,function(e){return n.contains(e,t)})}):1==this.length?n(T.qsa(this[0],t)):this.map(function(){return T.qsa(this,t)}):n()},closest:function(t,e){var i=this[0],r=!1;for("object"==typeof t&&(r=n(t));i&&!(r?r.indexOf(i)>=0:T.matches(i,t));)i=i!==e&&!$(i)&&i.parentNode;return n(i)},parents:function(t){for(var e=[],i=this;i.length>0;)i=n.map(i,function(t){return(t=t.parentNode)&&!$(t)&&e.indexOf(t)<0?(e.push(t),t):void 0});return U(e,t)},parent:function(t){return U(N(this.pluck("parentNode")),t)},children:function(t){return U(this.map(function(){return V(this)}),t)},contents:function(){return this.map(function(){return o.call(this.childNodes)})},siblings:function(t){return U(this.map(function(t,e){return s.call(V(e.parentNode),function(t){return t!==e})}),t)},empty:function(){return this.each(function(){this.innerHTML=""})},pluck:function(t){return n.map(this,function(e){return e[t]})},show:function(){return this.each(function(){"none"==this.style.display&&(this.style.display=""),"none"==getComputedStyle(this,"").getPropertyValue("display")&&(this.style.display=I(this.nodeName))})},replaceWith:function(t){return this.before(t).remove()},wrap:function(t){var e=Z(t);if(this[0]&&!e)var i=n(t).get(0),r=i.parentNode||this.length>1;return this.each(function(o){n(this).wrapAll(e?t.call(this,o):r?i.cloneNode(!0):i)})},wrapAll:function(t){if(this[0]){n(this[0]).before(t=n(t));for(var e;(e=t.children()).length;)t=e.first();n(t).append(this)}return this},wrapInner:function(t){var e=Z(t);return this.each(function(i){var r=n(this),o=r.contents(),s=e?t.call(this,i):t;o.length?o.wrapAll(s):r.append(s)})},unwrap:function(){return this.parent().each(function(){n(this).replaceWith(n(this).children())}),this},clone:function(){return this.map(function(){return this.cloneNode(!0)})},hide:function(){return this.css("display","none")},toggle:function(e){return this.each(function(){var i=n(this);(e===t?"none"==i.css("display"):e)?i.show():i.hide()})},prev:function(t){return n(this.pluck("previousElementSibling")).filter(t||"*")},next:function(t){return n(this.pluck("nextElementSibling")).filter(t||"*")},html:function(t){return 0 in arguments?this.each(function(e){var i=this.innerHTML;n(this).empty().append(J(this,t,e,i))}):0 in this?this[0].innerHTML:null},text:function(t){return 0 in arguments?this.each(function(e){var n=J(this,t,e,this.textContent);this.textContent=null==n?"":""+n}):0 in this?this[0].textContent:null},attr:function(n,i){var r;return"string"!=typeof n||1 in arguments?this.each(function(t){if(1===this.nodeType)if(D(n))for(e in n)X(this,e,n[e]);else X(this,n,J(this,i,t,this.getAttribute(n)))}):this.length&&1===this[0].nodeType?!(r=this[0].getAttribute(n))&&n in this[0]?this[0][n]:r:t},removeAttr:function(t){return this.each(function(){1===this.nodeType&&t.split(" ").forEach(function(t){X(this,t)},this)})},prop:function(t,e){return t=P[t]||t,1 in arguments?this.each(function(n){this[t]=J(this,e,n,this[t])}):this[0]&&this[0][t]},data:function(e,n){var i="data-"+e.replace(m,"-$1").toLowerCase(),r=1 in arguments?this.attr(i,n):this.attr(i);return null!==r?Y(r):t},val:function(t){return 0 in arguments?this.each(function(e){this.value=J(this,t,e,this.value)}):this[0]&&(this[0].multiple?n(this[0]).find("option").filter(function(){return this.selected}).pluck("value"):this[0].value)},offset:function(t){if(t)return this.each(function(e){var i=n(this),r=J(this,t,e,i.offset()),o=i.offsetParent().offset(),s={top:r.top-o.top,left:r.left-o.left};"static"==i.css("position")&&(s.position="relative"),i.css(s)});if(!this.length)return null;var e=this[0].getBoundingClientRect();return{left:e.left+window.pageXOffset,top:e.top+window.pageYOffset,width:Math.round(e.width),height:Math.round(e.height)}},css:function(t,i){if(arguments.length<2){var r,o=this[0];if(!o)return;if(r=getComputedStyle(o,""),"string"==typeof t)return o.style[C(t)]||r.getPropertyValue(t);if(A(t)){var s={};return n.each(t,function(t,e){s[e]=o.style[C(e)]||r.getPropertyValue(e)}),s}}var a="";if("string"==L(t))i||0===i?a=F(t)+":"+H(t,i):this.each(function(){this.style.removeProperty(F(t))});else for(e in t)t[e]||0===t[e]?a+=F(e)+":"+H(e,t[e])+";":this.each(function(){this.style.removeProperty(F(e))});return this.each(function(){this.style.cssText+=";"+a})},index:function(t){return t?this.indexOf(n(t)[0]):this.parent().children().indexOf(this[0])},hasClass:function(t){return t?r.some.call(this,function(t){return this.test(W(t))},q(t)):!1},addClass:function(t){return t?this.each(function(e){if("className"in this){i=[];var r=W(this),o=J(this,t,e,r);o.split(/\s+/g).forEach(function(t){n(this).hasClass(t)||i.push(t)},this),i.length&&W(this,r+(r?" ":"")+i.join(" "))}}):this},removeClass:function(e){return this.each(function(n){if("className"in this){if(e===t)return W(this,"");i=W(this),J(this,e,n,i).split(/\s+/g).forEach(function(t){i=i.replace(q(t)," ")}),W(this,i.trim())}})},toggleClass:function(e,i){return e?this.each(function(r){var o=n(this),s=J(this,e,r,W(this));s.split(/\s+/g).forEach(function(e){(i===t?!o.hasClass(e):i)?o.addClass(e):o.removeClass(e)})}):this},scrollTop:function(e){if(this.length){var n="scrollTop"in this[0];return e===t?n?this[0].scrollTop:this[0].pageYOffset:this.each(n?function(){this.scrollTop=e}:function(){this.scrollTo(this.scrollX,e)})}},scrollLeft:function(e){if(this.length){var n="scrollLeft"in this[0];return e===t?n?this[0].scrollLeft:this[0].pageXOffset:this.each(n?function(){this.scrollLeft=e}:function(){this.scrollTo(e,this.scrollY)})}},position:function(){if(this.length){var t=this[0],e=this.offsetParent(),i=this.offset(),r=d.test(e[0].nodeName)?{top:0,left:0}:e.offset();return i.top-=parseFloat(n(t).css("margin-top"))||0,i.left-=parseFloat(n(t).css("margin-left"))||0,r.top+=parseFloat(n(e[0]).css("border-top-width"))||0,r.left+=parseFloat(n(e[0]).css("border-left-width"))||0,{top:i.top-r.top,left:i.left-r.left}}},offsetParent:function(){return this.map(function(){for(var t=this.offsetParent||a.body;t&&!d.test(t.nodeName)&&"static"==n(t).css("position");)t=t.offsetParent;return t})}},n.fn.detach=n.fn.remove,["width","height"].forEach(function(e){var i=e.replace(/./,function(t){return t[0].toUpperCase()});n.fn[e]=function(r){var o,s=this[0];return r===t?_(s)?s["inner"+i]:$(s)?s.documentElement["scroll"+i]:(o=this.offset())&&o[e]:this.each(function(t){s=n(this),s.css(e,J(this,r,t,s[e]()))})}}),v.forEach(function(t,e){var i=e%2;n.fn[t]=function(){var t,o,r=n.map(arguments,function(e){return t=L(e),"object"==t||"array"==t||null==e?e:T.fragment(e)}),s=this.length>1;return r.length<1?this:this.each(function(t,u){o=i?u:u.parentNode,u=0==e?u.nextSibling:1==e?u.firstChild:2==e?u:null;var f=n.contains(a.documentElement,o);r.forEach(function(t){if(s)t=t.cloneNode(!0);else if(!o)return n(t).remove();o.insertBefore(t,u),f&&G(t,function(t){null==t.nodeName||"SCRIPT"!==t.nodeName.toUpperCase()||t.type&&"text/javascript"!==t.type||t.src||window.eval.call(window,t.innerHTML)})})})},n.fn[i?t+"To":"insert"+(e?"Before":"After")]=function(e){return n(e)[t](this),this}}),T.Z.prototype=n.fn,T.uniq=N,T.deserializeValue=Y,n.zepto=T,n}();window.Zepto=Zepto,void 0===window.$&&(window.$=Zepto),function(t){function l(t){return t._zid||(t._zid=e++)}function h(t,e,n,i){if(e=p(e),e.ns)var r=d(e.ns);return(s[l(t)]||[]).filter(function(t){return!(!t||e.e&&t.e!=e.e||e.ns&&!r.test(t.ns)||n&&l(t.fn)!==l(n)||i&&t.sel!=i)})}function p(t){var e=(""+t).split(".");return{e:e[0],ns:e.slice(1).sort().join(" ")}}function d(t){return new RegExp("(?:^| )"+t.replace(" "," .* ?")+"(?: |$)")}function m(t,e){return t.del&&!u&&t.e in f||!!e}function g(t){return c[t]||u&&f[t]||t}function v(e,i,r,o,a,u,f){var h=l(e),d=s[h]||(s[h]=[]);i.split(/\s/).forEach(function(i){if("ready"==i)return t(document).ready(r);var s=p(i);s.fn=r,s.sel=a,s.e in c&&(r=function(e){var n=e.relatedTarget;return!n||n!==this&&!t.contains(this,n)?s.fn.apply(this,arguments):void 0}),s.del=u;var l=u||r;s.proxy=function(t){if(t=j(t),!t.isImmediatePropagationStopped()){t.data=o;var i=l.apply(e,t._args==n?[t]:[t].concat(t._args));return i===!1&&(t.preventDefault(),t.stopPropagation()),i}},s.i=d.length,d.push(s),"addEventListener"in e&&e.addEventListener(g(s.e),s.proxy,m(s,f))})}function y(t,e,n,i,r){var o=l(t);(e||"").split(/\s/).forEach(function(e){h(t,e,n,i).forEach(function(e){delete s[o][e.i],"removeEventListener"in t&&t.removeEventListener(g(e.e),e.proxy,m(e,r))})})}function j(e,i){return(i||!e.isDefaultPrevented)&&(i||(i=e),t.each(E,function(t,n){var r=i[t];e[t]=function(){return this[n]=x,r&&r.apply(i,arguments)},e[n]=b}),(i.defaultPrevented!==n?i.defaultPrevented:"returnValue"in i?i.returnValue===!1:i.getPreventDefault&&i.getPreventDefault())&&(e.isDefaultPrevented=x)),e}function S(t){var e,i={originalEvent:t};for(e in t)w.test(e)||t[e]===n||(i[e]=t[e]);return j(i,t)}var n,e=1,i=Array.prototype.slice,r=t.isFunction,o=function(t){return"string"==typeof t},s={},a={},u="onfocusin"in window,f={focus:"focusin",blur:"focusout"},c={mouseenter:"mouseover",mouseleave:"mouseout"};a.click=a.mousedown=a.mouseup=a.mousemove="MouseEvents",t.event={add:v,remove:y},t.proxy=function(e,n){var s=2 in arguments&&i.call(arguments,2);if(r(e)){var a=function(){return e.apply(n,s?s.concat(i.call(arguments)):arguments)};return a._zid=l(e),a}if(o(n))return s?(s.unshift(e[n],e),t.proxy.apply(null,s)):t.proxy(e[n],e);throw new TypeError("expected function")},t.fn.bind=function(t,e,n){return this.on(t,e,n)},t.fn.unbind=function(t,e){return this.off(t,e)},t.fn.one=function(t,e,n,i){return this.on(t,e,n,i,1)};var x=function(){return!0},b=function(){return!1},w=/^([A-Z]|returnValue$|layer[XY]$)/,E={preventDefault:"isDefaultPrevented",stopImmediatePropagation:"isImmediatePropagationStopped",stopPropagation:"isPropagationStopped"};t.fn.delegate=function(t,e,n){return this.on(e,t,n)},t.fn.undelegate=function(t,e,n){return this.off(e,t,n)},t.fn.live=function(e,n){return t(document.body).delegate(this.selector,e,n),this},t.fn.die=function(e,n){return t(document.body).undelegate(this.selector,e,n),this},t.fn.on=function(e,s,a,u,f){var c,l,h=this;return e&&!o(e)?(t.each(e,function(t,e){h.on(t,s,a,e,f)}),h):(o(s)||r(u)||u===!1||(u=a,a=s,s=n),(r(a)||a===!1)&&(u=a,a=n),u===!1&&(u=b),h.each(function(n,r){f&&(c=function(t){return y(r,t.type,u),u.apply(this,arguments)}),s&&(l=function(e){var n,o=t(e.target).closest(s,r).get(0);return o&&o!==r?(n=t.extend(S(e),{currentTarget:o,liveFired:r}),(c||u).apply(o,[n].concat(i.call(arguments,1)))):void 0}),v(r,e,u,a,s,l||c)}))},t.fn.off=function(e,i,s){var a=this;return e&&!o(e)?(t.each(e,function(t,e){a.off(t,i,e)}),a):(o(i)||r(s)||s===!1||(s=i,i=n),s===!1&&(s=b),a.each(function(){y(this,e,s,i)}))},t.fn.trigger=function(e,n){return e=o(e)||t.isPlainObject(e)?t.Event(e):j(e),e._args=n,this.each(function(){e.type in f&&"function"==typeof this[e.type]?this[e.type]():"dispatchEvent"in this?this.dispatchEvent(e):t(this).triggerHandler(e,n)})},t.fn.triggerHandler=function(e,n){var i,r;return this.each(function(s,a){i=S(o(e)?t.Event(e):e),i._args=n,i.target=a,t.each(h(a,e.type||e),function(t,e){return r=e.proxy(i),i.isImmediatePropagationStopped()?!1:void 0})}),r},"focusin focusout focus blur load resize scroll unload click dblclick mousedown mouseup mousemove mouseover mouseout mouseenter mouseleave change select keydown keypress keyup error".split(" ").forEach(function(e){t.fn[e]=function(t){return 0 in arguments?this.bind(e,t):this.trigger(e)}}),t.Event=function(t,e){o(t)||(e=t,t=e.type);var n=document.createEvent(a[t]||"Events"),i=!0;if(e)for(var r in e)"bubbles"==r?i=!!e[r]:n[r]=e[r];return n.initEvent(t,i,!0),j(n)}}(Zepto),function(t){function h(e,n,i){var r=t.Event(n);return t(e).trigger(r,i),!r.isDefaultPrevented()}function p(t,e,i,r){return t.global?h(e||n,i,r):void 0}function d(e){e.global&&0===t.active++&&p(e,null,"ajaxStart")}function m(e){e.global&&!--t.active&&p(e,null,"ajaxStop")}function g(t,e){var n=e.context;return e.beforeSend.call(n,t,e)===!1||p(e,n,"ajaxBeforeSend",[t,e])===!1?!1:void p(e,n,"ajaxSend",[t,e])}function v(t,e,n,i){var r=n.context,o="success";n.success.call(r,t,o,e),i&&i.resolveWith(r,[t,o,e]),p(n,r,"ajaxSuccess",[e,n,t]),x(o,e,n)}function y(t,e,n,i,r){var o=i.context;i.error.call(o,n,e,t),r&&r.rejectWith(o,[n,e,t]),p(i,o,"ajaxError",[n,i,t||e]),x(e,n,i)}function x(t,e,n){var i=n.context;n.complete.call(i,e,t),p(n,i,"ajaxComplete",[e,n]),m(n)}function b(){}function w(t){return t&&(t=t.split(";",2)[0]),t&&(t==f?"html":t==u?"json":s.test(t)?"script":a.test(t)&&"xml")||"text"}function E(t,e){return""==e?t:(t+"&"+e).replace(/[&?]{1,2}/,"?")}function j(e){e.processData&&e.data&&"string"!=t.type(e.data)&&(e.data=t.param(e.data,e.traditional)),!e.data||e.type&&"GET"!=e.type.toUpperCase()||(e.url=E(e.url,e.data),e.data=void 0)}function S(e,n,i,r){return t.isFunction(n)&&(r=i,i=n,n=void 0),t.isFunction(i)||(r=i,i=void 0),{url:e,data:n,success:i,dataType:r}}function C(e,n,i,r){var o,s=t.isArray(n),a=t.isPlainObject(n);t.each(n,function(n,u){o=t.type(u),r&&(n=i?r:r+"["+(a||"object"==o||"array"==o?n:"")+"]"),!r&&s?e.add(u.name,u.value):"array"==o||!i&&"object"==o?C(e,u,i,n):e.add(n,u)})}var i,r,e=0,n=window.document,o=/<script\b[^<]*(?:(?!<\/script>)<[^<]*)*<\/script>/gi,s=/^(?:text|application)\/javascript/i,a=/^(?:text|application)\/xml/i,u="application/json",f="text/html",c=/^\s*$/,l=n.createElement("a");l.href=window.location.href,t.active=0,t.ajaxJSONP=function(i,r){if(!("type"in i))return t.ajax(i);var f,h,o=i.jsonpCallback,s=(t.isFunction(o)?o():o)||"jsonp"+ ++e,a=n.createElement("script"),u=window[s],c=function(e){t(a).triggerHandler("error",e||"abort")},l={abort:c};return r&&r.promise(l),t(a).on("load error",function(e,n){clearTimeout(h),t(a).off().remove(),"error"!=e.type&&f?v(f[0],l,i,r):y(null,n||"error",l,i,r),window[s]=u,f&&t.isFunction(u)&&u(f[0]),u=f=void 0}),g(l,i)===!1?(c("abort"),l):(window[s]=function(){f=arguments},a.src=i.url.replace(/\?(.+)=\?/,"?$1="+s),n.head.appendChild(a),i.timeout>0&&(h=setTimeout(function(){c("timeout")},i.timeout)),l)},t.ajaxSettings={type:"GET",beforeSend:b,success:b,error:b,complete:b,context:null,global:!0,xhr:function(){return new window.XMLHttpRequest},accepts:{script:"text/javascript, application/javascript, application/x-javascript",json:u,xml:"application/xml, text/xml",html:f,text:"text/plain"},crossDomain:!1,timeout:0,processData:!0,cache:!0},t.ajax=function(e){var a,o=t.extend({},e||{}),s=t.Deferred&&t.Deferred();for(i in t.ajaxSettings)void 0===o[i]&&(o[i]=t.ajaxSettings[i]);d(o),o.crossDomain||(a=n.createElement("a"),a.href=o.url,a.href=a.href,o.crossDomain=l.protocol+"//"+l.host!=a.protocol+"//"+a.host),o.url||(o.url=window.location.toString()),j(o);var u=o.dataType,f=/\?.+=\?/.test(o.url);if(f&&(u="jsonp"),o.cache!==!1&&(e&&e.cache===!0||"script"!=u&&"jsonp"!=u)||(o.url=E(o.url,"_="+Date.now())),"jsonp"==u)return f||(o.url=E(o.url,o.jsonp?o.jsonp+"=?":o.jsonp===!1?"":"callback=?")),t.ajaxJSONP(o,s);var C,h=o.accepts[u],p={},m=function(t,e){p[t.toLowerCase()]=[t,e]},x=/^([\w-]+:)\/\//.test(o.url)?RegExp.$1:window.location.protocol,S=o.xhr(),T=S.setRequestHeader;if(s&&s.promise(S),o.crossDomain||m("X-Requested-With","XMLHttpRequest"),m("Accept",h||"*/*"),(h=o.mimeType||h)&&(h.indexOf(",")>-1&&(h=h.split(",",2)[0]),S.overrideMimeType&&S.overrideMimeType(h)),(o.contentType||o.contentType!==!1&&o.data&&"GET"!=o.type.toUpperCase())&&m("Content-Type",o.contentType||"application/x-www-form-urlencoded"),o.headers)for(r in o.headers)m(r,o.headers[r]);if(S.setRequestHeader=m,S.onreadystatechange=function(){if(4==S.readyState){S.onreadystatechange=b,clearTimeout(C);var e,n=!1;if(S.status>=200&&S.status<300||304==S.status||0==S.status&&"file:"==x){u=u||w(o.mimeType||S.getResponseHeader("content-type")),e=S.responseText;try{"script"==u?(1,eval)(e):"xml"==u?e=S.responseXML:"json"==u&&(e=c.test(e)?null:t.parseJSON(e))}catch(i){n=i}n?y(n,"parsererror",S,o,s):v(e,S,o,s)}else y(S.statusText||null,S.status?"error":"abort",S,o,s)}},g(S,o)===!1)return S.abort(),y(null,"abort",S,o,s),S;if(o.xhrFields)for(r in o.xhrFields)S[r]=o.xhrFields[r];var N="async"in o?o.async:!0;S.open(o.type,o.url,N,o.username,o.password);for(r in p)T.apply(S,p[r]);return o.timeout>0&&(C=setTimeout(function(){S.onreadystatechange=b,S.abort(),y(null,"timeout",S,o,s)},o.timeout)),S.send(o.data?o.data:null),S},t.get=function(){return t.ajax(S.apply(null,arguments))},t.post=function(){var e=S.apply(null,arguments);return e.type="POST",t.ajax(e)},t.getJSON=function(){var e=S.apply(null,arguments);return e.dataType="json",t.ajax(e)},t.fn.load=function(e,n,i){if(!this.length)return this;var a,r=this,s=e.split(/\s/),u=S(e,n,i),f=u.success;return s.length>1&&(u.url=s[0],a=s[1]),u.success=function(e){r.html(a?t("<div>").html(e.replace(o,"")).find(a):e),f&&f.apply(r,arguments)},t.ajax(u),this};var T=encodeURIComponent;t.param=function(e,n){var i=[];return i.add=function(e,n){t.isFunction(n)&&(n=n()),null==n&&(n=""),this.push(T(e)+"="+T(n))},C(i,e,n),i.join("&").replace(/%20/g,"+")}}(Zepto),function(t){t.fn.serializeArray=function(){var e,n,i=[],r=function(t){return t.forEach?t.forEach(r):void i.push({name:e,value:t})};return this[0]&&t.each(this[0].elements,function(i,o){n=o.type,e=o.name,e&&"fieldset"!=o.nodeName.toLowerCase()&&!o.disabled&&"submit"!=n&&"reset"!=n&&"button"!=n&&"file"!=n&&("radio"!=n&&"checkbox"!=n||o.checked)&&r(t(o).val())}),i},t.fn.serialize=function(){var t=[];return this.serializeArray().forEach(function(e){t.push(encodeURIComponent(e.name)+"="+encodeURIComponent(e.value))}),t.join("&")},t.fn.submit=function(e){if(0 in arguments)this.bind("submit",e);else if(this.length){var n=t.Event("submit");this.eq(0).trigger(n),n.isDefaultPrevented()||this.get(0).submit()}return this}}(Zepto),function(t){"__proto__"in{}||t.extend(t.zepto,{Z:function(e,n){return e=e||[],t.extend(e,t.fn),e.selector=n||"",e.__Z=!0,e},isZ:function(e){return"array"===t.type(e)&&"__Z"in e}});try{getComputedStyle(void 0)}catch(e){var n=getComputedStyle;window.getComputedStyle=function(t){try{return n(t)}catch(e){return null}}}}(Zepto);
`)
var Side_menu_css = []byte(`
body {
	color: #777;
}

.pure-img-responsive {
	max-width: 100%;
	height: auto;
}

/*
 * Add transition to containers so they can push in and out.
 */
#layout,
#menu,
.menu-link {
	-webkit-transition: all 0.2s ease-out;
	-moz-transition: all 0.2s ease-out;
	-ms-transition: all 0.2s ease-out;
	-o-transition: all 0.2s ease-out;
	transition: all 0.2s ease-out;
}

#layout {
position: relative;
padding-left: 0;
}
#layout.active #menu {
left: 150px;
width: 150px;
}

#layout.active .menu-link {
left: 150px;
}
.content {
	margin: 0 auto;
	padding: 0 2em;
	max-width: 800px;
	margin-bottom: 50px;
	line-height: 1.6em;
}

.header {
	margin: 0;
	color: #333;
	text-align: center;
	padding: 2.5em 2em 0;
	border-bottom: 1px solid #eee;
}
.header h1 {
	margin: 0.2em 0;
	font-size: 3em;
	font-weight: 300;
}
.header h2 {
	font-weight: 300;
	color: #ccc;
	padding: 0;
	margin-top: 0;
}

.content-subhead {
	margin: 50px 0 20px 0;
	font-weight: 300;
	color: #888;
}

#menu {
margin-left: -150px; /* "#menu" width */
width: 150px;
position: fixed;
top: 0;
left: 0;
bottom: 0;
z-index: 1000; /* so the menu or its navicon stays above all content */
background: #191818;
overflow-y: auto;
-webkit-overflow-scrolling: touch;
}
/*
 *	All anchors inside the menu should be styled like this.
 */
#menu a {
color: #999;
border: none;
padding: 0.6em 0 0.6em 0.6em;
}

/*
 *	Remove all background/borders, since we are applying them to #menu.
 */
#menu .pure-menu,
#menu .pure-menu ul {
border: none;
background: transparent;
}

/*
 *	Add that light border to separate items into groups.
 */
#menu .pure-menu ul,
#menu .pure-menu .menu-item-divided {
border-top: 1px solid #333;
}
/*
 *		Change color of the anchor links on hover/focus.
 */
#menu .pure-menu li a:hover,
#menu .pure-menu li a:focus {
background: #333;
}

#menu .pure-menu-selected,
#menu .pure-menu-heading {
background: #1f8dd6;
}
#menu .pure-menu-selected a {
color: #fff;
}

/*
 *	This styles the menu heading.
 */
#menu .pure-menu-heading {
font-size: 110%;
color: #fff;
margin: 0;
}

/* -- Dynamic Button For Responsive Menu -------------------------------------*/

/*
 * The button to open/close the Menu is custom-made and not part of Pure. Here's
 * how it works:
 */

.menu-link {
	position: fixed;
	display: block; /* show this only on small screens */
	top: 0;
	left: 0; /* "#menu width" */
	background: #000;
	background: rgba(0,0,0,0.7);
	font-size: 10px; /* change this value to increase/decrease button size */
	z-index: 10;
	width: 2em;
	height: auto;
	padding: 2.1em 1.6em;
}

.menu-link:hover,
.menu-link:focus {
	background: #000;
}

.menu-link span {
	position: relative;
	display: block;
}

.menu-link span,
.menu-link span:before,
.menu-link span:after {
	background-color: #fff;
	width: 100%;
	height: 0.2em;
}

.menu-link span:before,
.menu-link span:after {
	position: absolute;
	margin-top: -0.6em;
	content: " ";
}

.menu-link span:after {
	margin-top: 0.6em;
}


/* -- Responsive Styles (Media Queries) ------------------------------------- */

@media (min-width: 48em) {
	
	.header,
	.content {
		padding-left: 2em;
		padding-right: 2em;
	}
	
	#layout {
	padding-left: 150px; /* left col width "#menu" */
	left: 0;
	}
	#menu {
	left: 150px;
	}
	
	.menu-link {
		position: fixed;
		left: 150px;
		display: none;
	}
	
	#layout.active .menu-link {
	left: 150px;
	}
}

@media (max-width: 48em) {
	/* Only apply this when the window is small. Otherwise, the following
	 *	case results in extra padding on the left:
	 * Make the window small.
	 * Tap the menu to trigger the active state.
	 * Make the window large again.
	 */
	#layout.active {
	position: relative;
	left: 150px;
	}
}
`)

var Ui_js = []byte(`
(function (window, document) {
	var layout   = document.getElementById('layout'),
		menu     = document.getElementById('menu'),
	    menuLink = document.getElementById('menuLink');

	function toggleClass(element, className) {
		var classes = element.className.split(/\s+/),
            length = classes.length,
		    i = 0;

		for(; i < length; i++) {
			if (classes[i] === className) {
				classes.splice(i, 1);
				break;
			}
		}
		// The className is not found
		if (length === classes.length) {
			classes.push(className);
		}
		element.className = classes.join(' ');
	}

	menuLink.onclick = function (e) {
		var active = 'active';
		
		e.preventDefault();
		toggleClass(layout, active);
		toggleClass(menu, active);
		toggleClass(menuLink, active);
	};

}(this, this.document));
`)
