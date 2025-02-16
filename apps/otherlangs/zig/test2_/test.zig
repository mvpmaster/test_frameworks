const std = @import("std");
const net = std.net;
const mem = std.mem;
const fs = std.fs;
const io = std.io;

pub fn main() anyerror!void {
    var args = std.process.args();
    const exe_name = args.next() orelse "zelf-zerve";
    const public_path = args.next() orelse {
        std.log.err("Usage: {s} <dir to serve files from>", .{exe_name});
        return;
    };

    var dir = try fs.cwd().openDir(public_path, .{});
    const self_addr = try net.Address.resolveIp("127.0.0.1", 9000);
    var listener = net.StreamServer.init(.{});
    try (&listener).listen(self_addr);

    std.log.info("Listening on {}; press Ctrl-C to exit...", .{self_addr});

    while ((&listener).accept()) |conn| {
        std.log.info("Accepted Connection from: {}", .{conn.address});

        serveFile(&conn.stream, dir) catch |err| {
            if (@errorReturnTrace()) |bt| {
                std.log.err("Failed to serve client: {}: {}", .{err, bt});
            } else {
                std.log.err("Failed to serve client: {}", .{err});
            }
        };

        conn.stream.close();
    } else |err| {
        return err;
    }
}
