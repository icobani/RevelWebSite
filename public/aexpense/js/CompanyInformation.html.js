/**
 * Created by ibrahimcobani on 31/01/2017.
 */

(function ($) {

  'use strict';

  //Date Pickers
  $(document).ready(function () {
    $('#Company\\.StartDateFiscalYearX').datepicker({
      language:"tr-TR",
      format: {
        toDisplay: function (date, format, language) {
          console.log(' - To Display ');
          var m = moment(date);
          return m.format('DD.MM.YYYY');
        },
        toValue: function (date, format, language) {
          console.log(date);
          var m = moment(date,'DD.MM.YYYY');
          return m.format('YYYY.MM.DD');
        }
      }
    }).on('changeDate', function(chosen_date) {
      var mDate = new moment(chosen_date.date);
      $('#Company\\.StartDateFiscalYear').val(mDate.format("YYYY-MM-DD"));
      console.log(mDate.format("YYYY-MM-DD"))
    });
  });

})(window.jQuery);