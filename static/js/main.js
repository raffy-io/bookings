document.addEventListener("alpine:init", () => {
  // 1. Your existing notification component
  Alpine.data("notification", () => ({
    init() {
      const Toast = Swal.mixin({
        toast: true,
        position: "top-end",
        showConfirmButton: false,
        timer: 3000,
      });

      Toast.fire({
        icon: this.$el.dataset.icon,
        title: this.$el.dataset.msg,
      });
    },
  }));

  // notie
  Alpine.data("notie", () => ({
    init() {
      // Read clean, un-interpolated data attributes from the DOM element
      const type = this.$el.dataset.type;
      const message = this.$el.dataset.msg;

      notie.alert({
        type: type || "info",
        text: message,
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
