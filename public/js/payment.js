document.addEventListener('alpine:init', () => {
    Alpine.data('paymentHandler', () => ({
		stripe: null,
		elements: null,	
        async initPayment(event) {
            const response = event.detail.xhr.responseText;
            this.stripe = Stripe(JSON.parse(response).publishableKey);
			const {clientSecret} = await fetch("/api/v0/payments/intent", { 
				method: "POST",
				headers: {
					"Content-Type": "application/json"
				},
				body: JSON.stringify({
					total: document.getElementById('cartItems').getAttribute('alert-data'),
				}),
			}).then(r => r.json())
		
			this.elements = this.stripe.elements({ clientSecret })
			this.elements.create('payment').mount('#paymentElement')
        },

		async confirmPayment() {
			const {error} = await this.stripe.confirmPayment({
				elements: this.elements, 
				confirmParams: {
					return_url: window.location.href.split('?')[0] + '/tracking'
				}
			})

			if (error) {
				window.toast(error.message, 500);
			}
		},
    }));
});
