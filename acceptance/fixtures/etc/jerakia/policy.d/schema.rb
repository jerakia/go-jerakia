policy :schema do
  lookup :default do
    datasource :file, {
      :docroot    => "/var/lib/jerakia/data",
      :format     => :yaml,
      :searchpath => [
        "host/#{scope[:hostname]}",
        "env/#{scope[:env]}",
        "common",
    ],
    }
  end
end

