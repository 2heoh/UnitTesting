use Mojo::Base -strict;

use Test::More;
use Test::Mojo;

my $t = Test::Mojo->new('Gazetiere');
$t->app->model(bless {}, 'Gazetiere::MockModel');

$t->get_ok('/')->status_is(200)->content_like(qr/country/i);
done_testing();

package Gazetiere::MockModel;

sub country_list {
    return {country => 'yes'}
}

1;