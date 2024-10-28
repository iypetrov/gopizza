document.addEventListener('alpine:init', () => {
    Alpine.data('paymentHandler', () => ({
		stripe: null,
		clientSecret: null,
		intentId: null,
		elements: null,	

		async initPayment(event) {
		    const response = event.detail.xhr.responseText;
		    const { publishableKey } = JSON.parse(response);
			
		    this.stripe = Stripe(publishableKey);
			
		    const paymentResponse = await fetch("/api/v0/payments/metadata", { 
		        method: "POST",
		        headers: {
		            "Content-Type": "application/json"
		        },
		        body: JSON.stringify({
		            email: document.getElementById('cartItems').getAttribute('alert-email'),
		            total: document.getElementById('cartItems').getAttribute('alert-total'),
		        }),
		    });
		
		    const { intentId, clientSecret } = await paymentResponse.json();
		
		    if (!clientSecret) {
		        console.error("Failed to retrieve clientSecret.");
		        return;
		    }
		
		    this.intentId = intentId;
		    this.clientSecret = clientSecret;
		
		    this.elements = this.stripe.elements({ clientSecret });
		    const paymentElement = this.elements.create('payment');
		    paymentElement.mount('#paymentElement');
		},

		async confirmPayment() {
	    try {
	        await fetch("/api/v0/orders", { 
	            method: "POST",
	            headers: {
	                "Content-Type": "application/json"
	            },
	            body: JSON.stringify({
	                intentId: this.intentId,
	                total: document.getElementById('cartItems').getAttribute('alert-total'),
	            }),
	        });

	        const { error } = await this.stripe.confirmPayment({
	            elements: this.elements, 
	            confirmParams: {
	                return_url: window.location.href.split('?')[0] + '/tracking'
	            }
	        });

	        if (error) {
	            window.toast(error.message, 500);
	        }
	    } catch (err) {
	        console.error("Error during payment confirmation:", err);
	        window.toast("Error processing payment. Please try again.", 500);
	    }
	},
    }));
});
