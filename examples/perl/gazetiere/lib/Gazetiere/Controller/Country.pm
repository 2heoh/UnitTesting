package Gazetiere::Controller::Country;
use Mojo::Base 'Mojolicious::Controller';

use Mojo::JSON qw(j);
use Data::Dumper;

# This action will render a template
sub list {
  my $self = shift;

  my $countries = $self->app->dbconn->run(fixup => sub {
      $_->selectall_hashref("SELECT * FROM address WHERE type='country'", "id");
  });

  $self->render(json => $countries);
}

1;
