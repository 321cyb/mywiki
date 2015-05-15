<!DOCTYPE html>
<html>
<head>
<title>Markdown Viewer</title>
<link rel="stylesheet" type="text/css" href="/_static/github-markdown.css">
<link rel="stylesheet" type="text/css" href="/_static/md.css">
<link rel="stylesheet" type="text/css" href="/_static/common.css">
<link rel="stylesheet" type="text/css" href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/8.4/styles/default.min.css">
</head>
<body class="markdown-body">

{{ content | safe }}

<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/2.1.3/jquery.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/8.4/highlight.min.js"></script>
<script>hljs.initHighlightingOnLoad();</script>
<script type="text/javascript" src="/_static/hilitor-utf8.js"></script>
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

</body>
</html>

