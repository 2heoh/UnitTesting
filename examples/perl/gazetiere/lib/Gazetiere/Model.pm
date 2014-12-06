package Gazetiere::Model;

sub new {
    my ($class, @args) = @_;

    my $dbconn = DBIx::Connector->new(@args);

    return bless { dbh => $dbconn }, $class;
}

sub country_list {
    my $self = shift;

    return $self->{dbh}->run(fixup => sub {
      $_->selectall_hashref("SELECT * FROM address WHERE type='country'", "id", {Slice => {}});
    });
}

1;