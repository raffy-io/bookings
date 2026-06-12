document.addEventListener("alpine:init", () => {
  // 1. Keep the standard component initialization (e.g., for Go template flash messages)
  Alpine.data(
    "notification",
    (initialMessage = "", initialIcon = "success") => ({
      init() {
        if (initialMessage) {
          this.showToast(initialMessage, initialIcon);
        }
      },
      showToast(message, iconType) {
        const Toast = Swal.mixin({
          toast: true,
          position: "top-end",
          showConfirmButton: false,
          timer: 3000,
        });

        Toast.fire({
          icon: iconType,
          title: message,
        });
      },
    }),
  );
});
