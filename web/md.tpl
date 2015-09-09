<!DOCTYPE html>
<html>
  <head>
    <title>{{title}}</title>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no"/>

    <link href="_static/css/bootstrap-custom.min.css" rel="stylesheet">
    <link href="_static/css/font-awesome-4.0.3.css" rel="stylesheet">
    <link rel="stylesheet" href="_static/css/highlight.css">
    <link href="_static/css/base.css" rel="stylesheet">
    <link href="_static/css/extra.css" rel="stylesheet">
    <link href="_static/css/md-toc.css" rel="stylesheet">

    <link href="_static/css/search-box.css" rel="stylesheet">
    <link href="_static/css/awesomplete.css" rel="stylesheet">

    <link rel="stylesheet" href="_static/css/responsivemultimenu.css" type="text/css"/>
    <script src="_static/js/jquery-1.10.2.min.js"></script>
    <script type="text/javascript" src="_static/js/responsivemultimenu.js"></script>
  </head>
<body>
        <div class="row">
                <div class="rmm style site-nav">
                    {{navTree | safe}}
                </div>
            <!--div class="col-md-3">
                <span class="search"></span>
            </div-->
        </div>

        <div class="container">
            <div class="col-md-9" role="main">
                {{content | safe}}
            </div>

            <div class="col-md-3 md-toc">
                {{toc | safe}}
            </div>
        </div>


<div id="sb-search" class="sb-search">
    <form method="GET" action="/_search">
        <input class="sb-search-input" placeholder="Enter your search term..." type="search" name="q" id="search">
        <input class="sb-search-submit" type="submit" value="">
        <span class="sb-icon-search"></span>
    </form>
</div>


<footer class="col-md-12">
    <hr>
    <center>Copyright 321cyb</center>
</footer>


<script src="_static/js/bootstrap-3.0.3.min.js"></script>
<script src="_static/js/highlight.pack.js"></script>
<script src="_static/js/base.js"></script>

<script type="text/javascript" src="/_static/js/hilitor-utf8.js"></script>
<script type="text/javascript">
  var getQueryVariable = function (variable) {
      var query = window.location.search.substring(1);
      var vars = query.split('&');
      for (var i = 0; i < vars.length; i++) {
          var pair = vars[i].split('=');
          if (decodeURIComponent(pair[0]) == variable) {
              return decodeURIComponent(pair[1]);
          }
      }
      console.log('Query variable %s not found', variable);
  }

  var myHilitor;

  $(function() {
    myHilitor = new Hilitor2();
    myHilitor.setMatchType("left");
    myHilitor.apply(getQueryVariable("q"));
  });
</script>


<!-- Added for search box and search box auto complete -->
<script src="_static/js/awesomplete.js"></script>
<script src="_static/js/search-box-classie.js"></script>
<script src="_static/js/search-box.js"></script>
<script type="text/javascript">
    new UISearch( document.getElementById( 'sb-search' ) );

    // get DOM element
    var input_ele = $('#sb-search > form > input.sb-search-input')[0];
    var autocomp = new Awesomplete(input_ele);
    autocomp.list = {{autoComplete | safe}};
</script>

<!-- currently using http://mkdocs.org/ cosmo style. -->

</body>
</html>
