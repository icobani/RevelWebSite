/**
 * Created by ibrahimcobani on 11/02/2017.
 */

(function ($) {
  'use strict';

  var originalVal = $.fn.val;
  $.fn.val = function () {
    var prev;
    if (arguments.length > 0) {
      prev = originalVal.apply(this, []);
    }
    var result = originalVal.apply(this, arguments);
    if (arguments.length > 0 && prev != originalVal.apply(this, []))
      $(this).change();
    return result;
  };


  //Autonumeric plug-in - automatic addition of dollar signs,etc controlled by tag attributes
  $('.autonumeric').autoNumeric('init');

  $(document).ready(function () {
    $("#ExpenseCategory\\.IsPublic").change(function () {
      console.log($("#ExpenseCategory\\.IsPublic").is(':checked'));
      $("#ExpenseCategory\\.IsPublic").val($("#ExpenseCategory\\.IsPublic").is(':checked'));
    });

    $("#ExpenseCategory\\.ReceiptsCheck").change(function () {
      console.log($("#ExpenseCategory\\.ReceiptsCheck").is(':checked'));
      $("#ExpenseCategory\\.ReceiptsCheck").val($("#ExpenseCategory\\.ReceiptsCheck").is(':checked'));
    });

    $("#ExpenseCategory\\.ProjectCheck").change(function () {
      console.log($("#ExpenseCategory\\.ProjectCheck").is(':checked'));
      $("#ExpenseCategory\\.ProjectCheck").val($("#ExpenseCategory\\.ProjectCheck").is(':checked'));
    });

    $("#ExpenseCategory\\.CommentsCheck").change(function () {
      console.log($("#ExpenseCategory\\.CommentsCheck").is(':checked'));
      $("#ExpenseCategory\\.CommentsCheck").val($("#ExpenseCategory\\.CommentsCheck").is(':checked'));
    });


  });

})(window.jQuery);