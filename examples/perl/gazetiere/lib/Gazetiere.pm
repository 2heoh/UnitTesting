package Gazetiere;
use Mojo::Base 'Mojolicious';
use DBIx::Connector;
use Gazetiere::Model;

has 'model' => sub {
    my $self = shift;

    my $model = Gazetiere::Model->new(
        'dbi:Pg:dbname=address;host=osm-db-dev.srv.pv.km;port=5432"', 
        'reader', 
        'reader', 
        {pg_utf8_strings => 0, RaiseError => 1, AutoCommit => 0}         
    );

    return $model;
};

# This method will run once at server start
sub startup {
  my $self = shift;

  # Documentation browser under "/perldoc"
  $self->plugin('PODRenderer');

  # Router
  my $r = $self->routes;

  # Normal route to controller
  $r->get('/')->to('country#list');
}

1;
