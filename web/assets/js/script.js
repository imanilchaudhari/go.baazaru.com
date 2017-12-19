/*	Table OF Contents
 ==========================
 Carousel
 Ajax Tab
 List view , Grid view  and compact view
 Global Plugins
 Customs Script
 Responsive cat-collapse for homepage
 */


$(document).ready(function () {



    var isMobile = function () {
        return /(iphone|ipod|ipad|android|blackberry|windows ce|palm|symbian)/i.test(navigator.userAgent);
    };

    var navbarSite = $('.navbar-site');

    // check if RTL or LTR
    var ifrtl = false;
    var dir = $("html").attr("dir");
    if(dir == "rtl")  {
        var ifrtl = true;
    }

    /*==================================
     Carousel
     ====================================*/

    // Featured Listings  carousel || HOME PAGE
    var owlitem = $(".item-carousel");

    owlitem.owlCarousel({
        //navigation : true, // Show next and prev buttons
        navigation: false,
        pagination: true,
        items: 5,
        itemsDesktopSmall: [979, 3],
        itemsTablet: [768, 3],
        itemsTabletSmall: [660, 2],
        itemsMobile: [400, 1]

    });

    // Custom Navigation Events
    $("#nextItem").click(function () {
        owlitem.trigger('owl.next');
    });
    $("#prevItem").click(function () {
        owlitem.trigger('owl.prev');
    });


    // Featured Listings  carousel || HOME PAGE
    var featuredListSlider = $(".featured-list-slider");

    featuredListSlider.owlCarousel({
        //navigation : true, // Show next and prev buttons
        navigation: false,
        pagination: false,
        items: 5,
        itemsDesktopSmall: [979, 3],
        itemsTablet: [768, 3],
        itemsTabletSmall: [660, 2],
        itemsMobile: [400, 1]


    });

    // Custom Navigation Events
    $(".featured-list-row .next").click(function () {
        featuredListSlider.trigger('owl.next');
    });

    $(".featured-list-row .prev").click(function () {
        featuredListSlider.trigger('owl.prev');
    });


    /*==================================
     Ajax Tab || CATEGORY PAGE
     ====================================*/


    $("#ajaxTabs li > a").click(function () {

        $("#allAds").empty().append("<div id='loading text-center'> <br> <img class='center-block' src='images/loading.gif' alt='Loading' /> <br> </div>");
        $("#ajaxTabs li").removeClass('active');
        $(this).parent('li').addClass('active');
        $.ajax({
            url: this.href, success: function (html) {
                $("#allAds").empty().append(html);
                $('.tooltipHere').tooltip('hide');
            }
        });
        return false;
    });

    urls = $('#ajaxTabs li:first-child a').attr("href");
    //alert(urls);
    $("#allAds").empty().append("<div id='loading text-center'> <br> <img class='center-block' src='images/loading.gif' alt='Loading' /> <br>  </div>");
    $.ajax({
        url: urls, success: function (html) {
            $("#allAds").empty().append(html);
            $('.tooltipHere').tooltip('hide');

            // default grid view class invoke into ajax content (product item)
            $(function () {
                $('.hasGridView .item-list').addClass('make-grid');
                $('.hasGridView .item-list').matchHeight();
                $.fn.matchHeight._apply('.hasGridView .item-list');
            });
        }
    });


    /*==================================
     List view clickable || CATEGORY
     ====================================*/

    // List view , Grid view  and compact view

    //  var selector doesn't work on ajax tab category.hhml. This variables elements disable for V1.6
    // var listItem = $('.item-list');
    // var addDescBox = $('.item-list .add-desc-box');
    //  var addsWrapper = $('.adds-wrapper');

    $('.list-view,#ajaxTabs li a').click(function (e) { //use a class, since your ID gets mangled
        e.preventDefault();
        $('.grid-view,.compact-view').removeClass("active");
        $('.list-view').addClass("active");
        $('.item-list').addClass("make-list").removeClass("make-grid make-compact");


        if ($('.adds-wrapper').hasClass('property-list')) {
            $('.item-list .add-desc-box').removeClass("col-md-9").addClass("col-md-6");
        } else {
            $('.item-list .add-desc-box').removeClass("col-md-9").addClass("col-md-7");
        }

        $(function () {
            $('.item-list').matchHeight('remove');
        });
    });

    $('.grid-view').click(function (e) { //use a class, since your ID gets mangled
        e.preventDefault();
        $('.list-view,.compact-view').removeClass("active");
        $(this).addClass("active");
        $('.item-list').addClass("make-grid").removeClass("make-list make-compact");


        if ($('.adds-wrapper').hasClass('property-list')) {
            // keep it for future
        } else {
            //
        }

        $(function () {
            $('.item-list').matchHeight();
            $.fn.matchHeight._apply('.item-list');
        });

    });

    $(function () {
        $('.hasGridView .item-list').matchHeight();
        $.fn.matchHeight._apply('.hasGridView .item-list');
    });

    $(function () {
        $('.row-featured .f-category').matchHeight();
        $.fn.matchHeight._apply('.row-featured .f-category');
    });

    $(function () {
        $('.has-equal-div > div').matchHeight();
        $.fn.matchHeight._apply('.row-featured .f-category');
    });


    $('.compact-view').click(function (e) { //use a class, since your ID gets mangled
        e.preventDefault();
        $('.list-view,.grid-view').removeClass("active");
        $(this).addClass("active");
        $('.item-list').addClass("make-compact").removeClass("make-list make-grid");


        if ($('.adds-wrapper').hasClass('property-list')) {
            $('.item-list .add-desc-box').addClass("col-md-9").removeClass('col-md-6');
        } else {
            $('.item-list .add-desc-box').toggleClass("col-md-9 col-md-7");
        }

        $(function () {
            $('.adds-wrapper .item-list').matchHeight('remove');
        });
    });


    /*==================================
     Global Plugins ||
     ====================================*/

    $('.long-list').hideMaxListItems({
        'max': 8,
        'speed': 500,
        'moreText': 'View More ([COUNT])'
    });


    $('.long-list-user').hideMaxListItems({
        'max': 12,
        'speed': 500,
        'moreText': 'View More ([COUNT])'
    });

    $(function () {
        $('[data-toggle="tooltip"]').tooltip()
    });

    $("select.selecter").niceSelect({ // custom scroll select plugin
    });

    $(".selectpicker").niceSelect({ // category list Short by
        // customClass: "select-sort-by"
    });

    $(".scrollbar").niceScroll();  //  customs scroll plugin

    $(".loginForm").validate(); //  form validation


    // smooth scroll to the ID
    $(document).on('click', 'a.scrollto', function (event) {
        event.preventDefault();
        $('html, body').animate({
            scrollTop: $($.attr(this, 'href')).offset().top
        }, 500);
    });


    /*=======================================================================================
     cat-collapse Homepage Category Responsive view
     ========================================================================================*/
   // $('.collapse-box .collapse').collapse();

    var catCollapse = $('.cat-collapse');

    $(window).bind('resize load', function () {

        if ($(this).width() < 767) {
           catCollapse.collapse('hide');
           catCollapse.on('show.bs.collapse', function () {
                $(this).prev('.cat-title').find('.icon-down-open-big').addClass("active-panel");
                //$(this).prev('.cat-title').find('.icon-down-open-big').toggleClass('icon-down-open-big icon-up-open-big');
            });

            catCollapse.on('hide.bs.collapse', function () {
                $(this).prev('.cat-title').find('.icon-down-open-big').removeClass("active-panel");
            })

        } else {
            $('#bd-docs-nav').collapse('show');
            catCollapse.collapse('show');
        }

    });

    // DEMO PREVIEW

    $(".tbtn").click(function () {
        $('.themeControll').toggleClass('active')
    });

    // jobs

    $("input:radio").click(function () {
        if ($('input:radio#job-seeker:checked').length > 0) {
            $('.forJobSeeker').removeClass('hide');
            $('.forJobFinder').addClass('hide');
        } else {
            $('.forJobFinder').removeClass('hide');
            $('.forJobSeeker').addClass('hide')
        }
    });

    $(".filter-toggle").click(function () {
        $('.mobile-filter-sidebar')
            .prepend("<div class='closeFilter'>X</div>")
            .animate({"left": "0"}, 250, "linear", function () {
        });
        $('.menu-overly-mask').addClass('is-visible');
    });

    $(".menu-overly-mask").click(function () {
        $(".mobile-filter-sidebar").animate({"left": "-251px"}, 250, "linear", function () {
        });
        $('.menu-overly-mask').removeClass('is-visible');
    });


    $(document).on('click', '.closeFilter', function () {
        $(".mobile-filter-sidebar").animate({"left": "-251px"}, 250, "linear", function () {
        });
        $('.menu-overly-mask').removeClass('is-visible');
    });


    // cityName will replace with selected location/area from location modal

    $('#selectRegion').on('shown.bs.modal', function (e) {
        // alert('Modal is successfully shown!');
        $("ul.list-link li a").click(function () {
            $('ul.list-link li a').removeClass('active');
            $(this).addClass('active');
            $(".cityName").text($(this).text());
            $('#selectRegion').modal('hide');
        });
    });

    $("#checkAll").click(function () {
        $('.add-img-selector input:checkbox').not(this).prop('checked', this.checked);
    });


    var stickyScroller = function () {
        var intialscroll = 0;
        $(window).scroll(function (event) {
            var windowScroll = $(this).scrollTop();
            if (windowScroll > intialscroll) {
                // downward-scrolling
                navbarSite.addClass('stuck');
            } else {
                // upward-scrolling
                navbarSite.removeClass('stuck');
            }
            if (windowScroll < 450) {
                // downward-scrolling
                navbarSite.removeClass('stuck');
            }
            intialscroll = windowScroll;
        });
    };

    if (isMobile()) {
        // For  mobile , ipad, tab
    } else {
        stickyScroller();
    } // end Desktop else

}); // end Ready


	
