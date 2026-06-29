document.addEventListener("alpine:init", () => {
  // notie
  Alpine.data("notie", () => ({
    init() {
      notie.alert({
        type: this.$el.dataset.type,
        text: this.$el.dataset.msg,
      });
    },
  }));

  // Datepicker
  Alpine.data("datePickerComponent", () => ({
    picker: null,

    init() {
      this.picker = new DateRangePicker(this.$refs.pickerElement, {
        minDate: new Date(),
        autohide: true,
        todayHighlight: true,
        format: "yyyy-mm-dd",
      });
    },
  }));
});
