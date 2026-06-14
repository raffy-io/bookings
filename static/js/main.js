document.addEventListener("alpine:init", () => {
  // Swal2 Toast
  Alpine.data("notification", () => ({
    init() {
      const Toast = Swal.mixin({
        toast: true,
        position: "top-end",
        showConfirmButton: false,
        timer: 3000,
        timerProgressBar: true,
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

  // AjaxModal
  Alpine.data("ajaxModal", () => ({
    async availabilitySearch() {
      const form = this.$refs.availabilityAjaxForm;
      const formData = new FormData(form);
      formData.append("csrf_token", "csrfToken");

      try {
        const res = await fetch("/search-availability-json", {
          method: "POST",
          body: formData,
        });

        if (!res.ok) {
          throw new Error(`HTTP error! status: ${res.status}`);
        }

        const data = await res.json();
        console.log(data);
      } catch (error) {
        console.error("Failed to fetch availability:", error);
      }
    },
  }));
});
