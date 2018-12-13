policy :autorun do

  lookup :default do
    datasource :file, {
      :format => :yaml,
      :docroot => "/var/lib/hiera",
      :map_namespace => false,
      :searchpath => [
        "common"
      ],
    }
  end
end

