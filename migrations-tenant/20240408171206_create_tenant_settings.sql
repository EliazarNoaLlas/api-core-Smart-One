-- +goose Up
-- +goose StatementBegin
create table if not exists db_tenant.tenant_settings
(
    id         varchar(36)  not null
        primary key,
    tenant_id  varchar(36)  not null,
    code       varchar(100) not null,
    value      varchar(250) not null,
    type       varchar(20)  not null comment 'public or private',
    enable     tinyint(1)   not null,
    created_at datetime     not null,
    deleted_at datetime     null,
    constraint tenant_settings_tenants_x_tenant_id_fk
        foreign key (tenant_id) references db_tenant.tenants (x_tenant_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE db_tenant.tenant_settings;
-- +goose StatementEnd
