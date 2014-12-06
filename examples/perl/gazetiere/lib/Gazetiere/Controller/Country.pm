package Gazetiere::Controller::Country;

use Mojo::Base 'Mojolicious::Controller';
use Mojo::JSON qw(j);
use Encode;

sub pretty {
    my $countries = shift;

    my @out = ();
    foreach my $id (keys  %$countries) {
        # decode text to JSON
        $countries->{$id}{info} = j(Encode::encode('utf8',$countries->{$id}{info}))
            if exists $countries->{$id}{info};

        # delete empty fields
        foreach (keys %{$countries->{$id}}) {
            delete $countries->{$id}{$_}
                unless defined $countries->{$id}{$_};
        }

        push @out, $countries->{$id};
    }

    return \@out
}

# This action will render a template
sub list {
    my $self = shift;

    $self->render(
        json => pretty($self->app->model->country_list)
    );
}

1;
