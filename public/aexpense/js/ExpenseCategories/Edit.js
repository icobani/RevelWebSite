/**
 * Created by ibrahimcobani on 11/02/2017.
 */


/**
 * Şube değiştiği zaman departmanları doldurmak için kullanılacak.
 * @param myselect : şube kodu gelir.
 * @constructor
 */
function BranchChanged(myselect) {
  //alert(myselect.value);
  $('#ExpenseCategory\\.DepartmentId').empty();



  $.getJSON("/api/Departments/" + myselect.value, null, function(result) {
    $("#ExpenseCategory\\.DepartmentId option").remove(); // Remove all <option> child tags.
    $.each(result.data, function(index, item) { // Iterates through a collection
      $("#ExpenseCategory\\.DepartmentId").append( // Append an object to the inside of the select box
          $("<option></option>") // Yes you can do this.
              .text(item.Value)
              .val(item.Id)
      );
    });
  });


}

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
    $("#ExpenseCategory\\.ReceiptsCheck").change(function () {
      $("#ExpenseCategory\\.ReceiptsCheck").val($("#ExpenseCategory\\.ReceiptsCheck").is(':checked'));
    });

    $("#ExpenseCategory\\.ProjectCheck").change(function () {
      $("#ExpenseCategory\\.ProjectCheck").val($("#ExpenseCategory\\.ProjectCheck").is(':checked'));
    });

    $("#ExpenseCategory\\.CommentsCheck").change(function () {
      $("#ExpenseCategory\\.CommentsCheck").val($("#ExpenseCategory\\.CommentsCheck").is(':checked'));
    });
  });





})(window.jQuery);