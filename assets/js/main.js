new Vue({
    el: '#links-app',

    data: function() {
        return {
            links: []
        }
    },

    created: function() {
        this.fetchLinks();
    },

    methods: {
        fetchLinks: function() {
            this.$http.get('api/links', function(data) {
                this.links = data.links;
            }.bind(this));
        },

        deleteLink: function(link) {
            var resource = this.$resource('api/links/{id}');

            resource.delete({ id: link.ID }, function(data) {
                console.log(data);
            }.bind(this));
        }
    }
})
