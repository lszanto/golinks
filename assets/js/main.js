new Vue({
    el: '#links-app',

    http: {
        root: '/api',
        headers: {

        }
    },

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
            this.$http.get('links', function(data) {
                this.links = data.links;
            }.bind(this));
        },

        authoriseUser: function() {
            this.$http.post('user/login', { 'username': 'luke', 'password': 'cool' }, function(data) {
                alert('HELLO CAME BACK');
            });
        },

        deleteLink: function(link) {
            var resource = this.$resource('api/links/{id}');

            resource.delete({ id: link.ID }, function(data) {
                console.log(data);
            }.bind(this));
        }
    }
})
