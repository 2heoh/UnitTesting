package Gazetiere::Controller::Country;
use Mojo::Base 'Mojolicious::Controller';

use Mojo::JSON qw(j);
use Data::Dumper;

# This action will render a template
sub list {
  my $self = shift;

  $self->render(json => $self->app->model->country_list);
}

1;
