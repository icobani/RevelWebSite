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

  $(document).ready(function () {

  });

})(window.jQuery);