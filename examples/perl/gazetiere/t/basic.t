use Mojo::Base -strict;

use Test::More;
use Test::Mojo;

my $t = Test::Mojo->new('Gazetiere');


$t->get_ok('/')->status_is(200)->content_like(qr/country/i);
done_testing();


1;