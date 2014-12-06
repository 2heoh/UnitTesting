use Mojo::Base -strict;

use Test::More;
use Test::Mojo;

# Arrange
my $t = Test::Mojo->new('Gazetiere');
$t->app->model(bless {}, 'Gazetiere::FakeModel');

# Act
$t->get_ok('/')
# Assert
->status_is(200)
->content_like(qr/country/i);

done_testing();


package Gazetiere::FakeModel;

sub country_list { { 1 => {country => 'yes'} } } 