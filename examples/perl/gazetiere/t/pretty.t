use Mojo::Base -strict;

use Test::More tests => 3;
use Test::Mojo;

use Gazetiere::Controller::Country;

# Arrange
my %countries = (
    '1' => {
        type => 'county',
        info => '{}',
        name => 'test',
        empty_field => undef,
    }
); 
# Act
my $res = Gazetiere::Controller::Country::pretty(\%countries);
# Assert
ok(ref $res eq 'ARRAY', 'got array');
ok($res->[0]{name} eq 'test', 'name match');
ok(!exists $res->[0]{empty_field}, 'empty field omitted');

