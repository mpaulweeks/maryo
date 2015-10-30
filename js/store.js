
function index(){
	console.log('start');
	$.ajax({
		url: "https://miiverse.nintendo.net/users/MrLuckyWaffles/posts",
		type: "GET",
	}).success(function(res) {
		console.log(res);
		var out = res.results[0];
		// var out = $(res.responseText).find('a.tsh').text();
		$("#results").html(out);
  	});
}