import { showFeedback } from "./script.js";

const verificationForm = document.getElementById("verificationForm");
const verificationCodeInput = document.getElementById("verificationCode");
const verificationBtn = document.getElementById("verificationBtn");

verificationForm.addEventListener("submit", async (event) => {
  event.preventDefault();

  const verificationCode = verificationCodeInput.value.trim();

  if (!verificationCode) {
    showFeedback("Please enter a verification code.", false);
    return;
  }

  verificationBtn.disabled = true;
  verificationBtn.textContent = "Verifying...";

  try {
    const response = await fetch("/verify", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ verification_code: verificationCode }),
    });

    const data = await response.json();

    if (response.ok && data.success) {
      showFeedback(data.message, data.success);

      setTimeout(() => {
        window.location.href = "/sign-in";
      }, 1000);
    } else {
      const errorData = await response.json();
      showFeedback(errorData.message || errorData.error, false);
      verificationBtn.disabled = false;
      verificationBtn.textContent = "Verify";
      return;
    }
  } catch (error) {
    console.error(error);
    showFeedback("Failed to verify. Please try again.", false);
    verificationBtn.disabled = false;
    verificationBtn.textContent = "Verify";
    return;
  }
});
