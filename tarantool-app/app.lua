#!/usr/bin/env tarantool
box.cfg {
    listen = 3301,
    background = true,
    log = '1.log',
    pid_file = '1.pid'
}
box.once("bootstrap", function()
    box.schema.space.create('dialog')
    box.space.dialog:create_index('primary', {unique=true, type = 'TREE', parts = {1, 'unsigned'}})
end)