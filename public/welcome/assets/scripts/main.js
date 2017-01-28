$(function () {

  new WOW().init();


  $(window).scroll(function () {
    var top = 71

    if ($(".lang-note").length) {
      $(".header").css("transition-property", "none")
      top -= $(".lang-note").outerHeight()
    }

    var sTop = $(window).scrollTop()
    if (sTop > top) {

      $('.index .logo').html('<img src="public/welcome/assets/images/logo.png" alt=""/>');
      $('.index').css('background-color', '#fff');
      $('.index ul li a').css('color', '#a3a3a3');

      if ($(".lang-note").length) {
        if (sTop > $(".lang-note").outerHeight()) {
          $('.index').css('top', 0);
        } else {
          $('.index').css('top', $(".lang-note").outerHeight() - sTop);
        }
      }
    }
    else {

      $('.index .logo').html('<img src="public/welcome/assets/images/logo-white.png" alt=""/>');
      $('.index').css('background-color', 'transparent');
      $('.index ul li a').css('color', '#fff');
    }
  });

  $('.arrow-down').click(function () {
    $('html, body').animate({scrollTop: 613}, 'slow');
    return false;
  });

  $('.toggle-bar').click(function () {

    $('ul.menu').toggle(400);

  });


  $(".register-form").submit(function (e) {
    toastr.warning('Server Hatası');
    e.preventDefault()
    var emailSelector = $('input[type="email"]');
    var email = emailSelector.val();
    if (email === "") return;

    var t = $(this);
    t.find('button[type="submit"] span.text').css("display", "none");
    t.find('button[type="submit"] span.spinner').addClass("three-quarters-loader");

    $.ajax({
      url: '/register',
      data: {email: email},
      type: 'post',
    }).success(function () {
      location.href = "/"
    }).fail(function (a, b, c) {
      if (a.status === 400) {
        emailSelector.val("")
        toastr.error('The email address invalid or taken')
      } else {
        toastr.error('Server Hatası')
      }
      t.find('button[type="submit"] span.text').css("display", "block")
      t.find('button[type="submit"] span.spinner').removeClass("three-quarters-loader")
    })
  });

  $(".header .lang").click(function (ev) {
    ev.stopPropagation()
    var e = $(".lang-menu")
    e.css("display", e.css("display") === "none" ? "block" : "none")
  })

  $(document).click(function () {
    if ($(this).closest(".lang-menu").length === 0) {
      $(".lang-menu").css("display", "none")
    }
  })

});
